package main

import (
	"appliedConcurrency/colors"
	"appliedConcurrency/endpoints"
	"fmt"
)

func main() {
	fmt.Printf("----\n \t%s\n----\n\n", colors.SprintFColor(colors.Red, "Go Concurrency Course"))
	endpoints.PrepareAndStartServer()
}
