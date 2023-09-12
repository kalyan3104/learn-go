package main

import (
	"fmt"
)

func main() {
	in := make(chan string)
	out := make(chan string)
	go func() {
		name := <-in
		out <- fmt.Sprintf("Hello, " + name)
	}()
	in <- "Rashik"
	close(in)
	message := <-out
	fmt.Println(message)
}
