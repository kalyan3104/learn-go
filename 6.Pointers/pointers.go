package main

import "fmt"

func main() {
	// var x *int
	// fmt.Println(x)
	// fmt.Println(*x) // panic: runtime error: invalid memory address or nil pointer dereference

	x := new(int)
	fmt.Println(x)
	fmt.Println(*x)

	y := 20
	fmt.Println("Before setToTen() ::: ", y) // y = 20
	setToTen(&y)
	fmt.Println("After setToTen() ::: ", y) // y = 10

	a := 10
	b := &a
	c := a
	fmt.Println(a, b, *b, c)

	a = 20 // Value of b becomes 20 but not c(10)
	fmt.Println(a, b, *b, c)

	*b = 30 // Value of a becomes 30 but not c(10)
	fmt.Println(a, b, *b, c)

	c = 40 // Only value of c is changes but not a,b
	fmt.Println(a, b, *b, c)
}

func setToTen(a *int) {
	*a = 10
}
