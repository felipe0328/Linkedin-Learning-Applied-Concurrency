package main

import (
	"appliedConcurrency/basics"
	"appliedConcurrency/channels"
	"appliedConcurrency/colors"
	"appliedConcurrency/patterns"
	"fmt"
)

func main() {
	fmt.Printf("----\n \t%s\n----\n\n", colors.SprintFColor(colors.Red, "Go Concurrency Course"))
	basics.RunBasics()
	channels.RunChannels()
	patterns.RunPatterns()
}
