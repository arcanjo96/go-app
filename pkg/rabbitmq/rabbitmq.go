package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChannel() (*amqp.Channel, error) {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672")
	if err != nil {
		return nil, err
	}
	channel, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	return channel, nil
}

func Consume(channel *amqp.Channel, outputChannel chan amqp.Delivery) error {
	msgs, err := channel.Consume(
		"order",
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}
	for msg := range msgs {
		outputChannel <- msg
	}
	return nil
}
