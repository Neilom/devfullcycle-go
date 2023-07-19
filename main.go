package main

import (
	"api-go/internal/infra/database"
	"api-go/internal/usecase"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Car struct {
	Model string
	Color string
}

// Metodo
func (c Car) Start() {
	println(c.Model + " has been started")
}

// Função
func soma(x, y int) int {
	return x + y
}

func main() {
	db, err := sql.Open("sqlite3", "db.sqlite3")
	if err != nil {
		panic(err)
	}

	orderRepository := &database.OrderRepository{
		Db: db,
	}

	usecaseda := usecase.CalculateFinalPrice{
		OrderRepository: orderRepository,
	}

	input := usecase.OrderInput{
		ID:    "123",
		Price: 10,
		Tax:   1,
	}

	output, err := usecaseda.Execute(input)
	if err != nil {
		panic(err)
	}
	println(output)
}
