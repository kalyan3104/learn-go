// go run switch.go YourValue
package main

import (
	"fmt"
	"os"
)

func main() {
	word := os.Args[1]
	// if word == "hello" {
	// 	fmt.Println("Hi yourself")
	// } else if word == "goodbye" {
	// 	fmt.Println("Take care")
	// } else if word == "greetings" {
	// 	fmt.Println("Salutations!")
	// } else {
	// 	fmt.Println("I dont know what you said")
	// }
	switch word {
	case "hi":
		fmt.Println("Very informal")
		fallthrough
	case "hello":
		fmt.Println("Hi Yourself")
	case "goodbye", "bye": // Having more than one case
		fmt.Println("Take care")
	case "greetings":
		fmt.Println("Salutations!")
	default:
		fmt.Println("I don't know what you said")
	}
}
