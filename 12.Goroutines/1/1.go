package main

import "fmt"

func main() {
	go runMe()
}

func runMe() {
	fmt.Println("Hello from a Goroutine")
}
