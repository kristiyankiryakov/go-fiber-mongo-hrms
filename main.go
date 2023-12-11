package main

import(
	"github.com/kristiyankiryakov/go-fiber-mongo-hrms/db"
)

type Employee struct {
	ID     string
	Name   string
	Salary float64
	Age    float64
}

func main() {
	db.LoadDb()
}

