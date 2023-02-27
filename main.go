package main

import (
	"appliedConcurrency/colors"
	"appliedConcurrency/controller"
	"appliedConcurrency/db"
	"appliedConcurrency/endpoints"
	"fmt"
	"log"
)

func main() {
	fmt.Printf("----\n \t%s\n----\n\n", colors.SprintFColor(colors.Red, "Go Concurrency Course"))
	ordersDB := db.NewOrders()
	productDB, err := db.NewProducts()

	if err != nil {
		log.Fatal(err)
	}

	controller := controller.NewController(productDB, ordersDB)

	endpoints.PrepareAndStartServer(controller)
}
