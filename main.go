package main

import (
	"appliedConcurrency/colors"
	"fmt"
)

func main() {
	fmt.Printf("----\n \t%s\n----\n\n", colors.SprintFColor(colors.Red, "Go Concurrency Course"))
}
