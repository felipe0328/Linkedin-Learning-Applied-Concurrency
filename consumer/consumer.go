package main

import (
	"appliedConcurrency/colors"
	"appliedConcurrency/models"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const (
	serverUrl       = "https://felipe0328-obscure-invention-wx5xpx9g56375-3000.preview.app.github.dev"
	index           = "/"
	createOrder     = "/orders"
	maxOrderAmount  = 15
	simulationCount = 100
)

var products []string = []string{"MWBLU", "MWLEM", "MWORG", "MWPEA", "MWRAS", "MWSTR", "MWCRA", "MWMAN"}

func main() {
	colors.PrintlnColor(colors.Blue, "Welcome to the consumer simulator!")
	err := testConnection()
	if err != nil {
		log.Fatalf("Server is not running, %s:%s", colors.SprintFColor(colors.Red, "Error"), err)
	}

	fmt.Printf("%s: Successful\n", colors.SprintFColor(colors.Green, "Testing Connection Status"))

	var wg sync.WaitGroup
	wg.Add(simulationCount)
	for i := 0; i < simulationCount; i++ {
		go createNewOrder(i, &wg)
	}
	wg.Wait()
}

func createNewOrder(number int, wg *sync.WaitGroup) {
	defer wg.Done()
	rand.New(rand.NewSource(time.Now().UnixNano()))
	amount := rand.Intn((maxOrderAmount) + 1)
	product := products[rand.Intn(len(products))]

	item := models.Item{
		Amount:    amount,
		ProductID: product,
	}

	fmt.Printf("%s %d request: %+v\n", colors.SprintFColor(colors.Green, "Creating new request for requester:"), number, item)
	itemBytes, err := json.Marshal(item)
	if err != nil {
		log.Fatal(err)
	}

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", serverUrl, createOrder), bytes.NewBuffer(itemBytes))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s %d\n", colors.SprintFColor(colors.Green, "Request Completed for requester:"), number)
}

func testConnection() error {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", serverUrl, index), nil)
	if err != nil {
		return err
	}

	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}
