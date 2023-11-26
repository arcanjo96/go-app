package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	amqp "github.com/rabbitmq/amqp091-go"

	"github.com/arcanjo96/go-test/internal/infra/database"
	usecase "github.com/arcanjo96/go-test/internal/usecase"
	"github.com/arcanjo96/go-test/pkg/rabbitmq"
)

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)
	channel, err := rabbitmq.OpenChannel()
	if err != nil {
		panic(err)
	}
	defer channel.Close()
	msgRabbitMqChannel := make(chan amqp.Delivery)
	go rabbitmq.Consume(channel, msgRabbitMqChannel) // escutando fila
	rabbitmqWorker(msgRabbitMqChannel, uc)           // thread1
}

func rabbitmqWorker(msgChannel chan amqp.Delivery, uc *usecase.CalculateFinalPrice) {
	fmt.Println("Starting rabbitmq")
	for msg := range msgChannel {
		var input usecase.OrderInput
		err := json.Unmarshal(msg.Body, &input)
		if err != nil {
			panic(err)
		}
		output, err := uc.Execute(input)
		if err != nil {
			panic(err)
		}
		msg.Ack(false)
		fmt.Println("Mensagem processada e salva no banco: ", output)
	}
}
