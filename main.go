package main

import (
	"database/sql"
	//"fmt"
	"github.com/devfullcycle/go-intensivo-p/internal/infra/database"
	"github.com/devfullcycle/go-intensivo-p/internal/usecase"
	_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	model string
	color string
}

func (c Car) Start() {
	println(c.model + " is started")
}

func soma(x, y int) int {
	return x + y
}

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}
	orderRepository := database.NewOrderRepository(db)
	uc := usecase.NewCalculateFinalPrice(orderRepository)

	input := usecase.OrderInput{
		ID:    "123",
		Price: 10.0,
		Tax:   1.0,
	}
	output, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	println(output)
	//fmt.println(output)
	// order, err := entity.NewOrder("1", -10, 1)
	// if err != nil {
	// 	println(err.Error())
	// } else {
	// 	println(order.ID)
	// }
	// car := Car{ //Declarando e atribuindo var
	// 	model: "toyota",
	// 	color: "amarelo",
	// }
	// car.Start()
	// println(car.model, car.color)
}
