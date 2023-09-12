package main

import "fmt"

func main() {
	// Arrays length is a part of its type. Arrays cannot be resized.
	var vals [3]int
	vals[0] = 2
	vals[1] = 4
	vals[2] = 6
	fmt.Println(vals)

	x := [5]int{1, 2, 3, 4, 5}
	fmt.Println(x)

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a)

	b := [2]string{"firstName", "lastName"}
	fmt.Println(b)

}
