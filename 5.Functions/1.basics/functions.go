package main

import "fmt"

// Go does not support function overloading
// functions in go are called by value
//func FunctionName(arg type) returnType {}

func main() {
	a := addNumbers(2, 3)
	fmt.Println("Addition of 2 and 3 is ", a)

	b, c := divAndRemainder(6, 2)
	fmt.Println("The quoteint is", b, "and remainder is", c)

	// Example of call by value
	x := 2
	y := [2]int{4, 6}
	z := "hello"
	doubleFail(x, y, z)
	fmt.Println("in [main]:::", x, y, z)
}

func addNumbers(a int, b int) int {
	return a + b
}

func divAndRemainder(a int, b int) (int, int) {
	return a / b, a % b
}

func doubleFail(a int, arr [2]int, s string) {
	a = a * 2
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * 2
	}
	s = s + s
	fmt.Println("in [doubleFail]:::", a, arr, s)
}
