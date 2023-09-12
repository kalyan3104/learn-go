package main

import "fmt"

func main() {
	// Loop one
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("Even\t", i)
		} else {
			fmt.Println("Odd\t", i)
		}
	}

	// Loop two
	j := 0
	for j < 10 {
		fmt.Printf("%d\t", j)
		j++
	}
	fmt.Println(j) // var j can be accessible out of the loop bcoz it's declared outside of loop

	// Loop three
	s := "Hello, World!"
	for k, v := range s {
		fmt.Println(k, v, string(v))
	}
}
