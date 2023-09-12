package main

import (
	"fmt"
	"time"
)

func main() {
	go runMe()
	time.Sleep(1 * time.Second)
}

func runMe() {
	fmt.Println("Hello from a Goroutine")
}
