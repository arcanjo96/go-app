package main

import (
	"net/http"

	"github.com/arcanjo96/go-test/internal/entity"
	"github.com/labstack/echo/v4"
)

func main() {
	// example using Chi
	// router := chi.NewRouter()
	// router.Use(middleware.Logger)
	// router.Get("/order", OrderHandler)
	// http.ListenAndServe(":8888", router)
	server := echo.New()
	server.GET("/order", OrderHandler)
	server.Logger.Fatal(server.Start(":8888"))
}

func OrderHandler(context echo.Context) error {
	order, _ := entity.NewOrder("123", 10, 5)
	err := order.CalculateFinalPrice()
	if err != nil {
		return context.JSON(500, err)
	}
	return context.JSON(http.StatusOK, order)
}

// example using go http
// func OrderHandler(response http.ResponseWriter, request *http.Request) error {
// order, _ := entity.NewOrder("123", 10, 5)
// err := order.CalculateFinalPrice()
// if err != nil {
// 	response.WriteHeader(http.StatusInternalServerError)
// }
// json.NewEncoder(response).Encode(order)
// }
