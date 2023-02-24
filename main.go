package main

import (
	"appliedConcurrency/basics"
	"appliedConcurrency/channels"
	"appliedConcurrency/patterns"
	"fmt"
)

func main() {
	fmt.Printf("----\n \tGo Concurrency Course\n----\n")
	basics.RunBasics()
	channels.RunChannels()
	patterns.RunPatterns()
}
