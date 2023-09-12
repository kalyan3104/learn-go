package main

import "fmt"

func main() {
	// Function inside of another function
	b := 2
	myAddTwo := func(a int) int {
		// return a + 2
		return a + b
	}
	fmt.Println("Function with in function::: 2+2 =", myAddTwo(4))

	printOperation(1, addOne)
	printOperation(2, addTwo)

	addingOne := makeAdder(1)
	fmt.Println("Function returning another function::: ", addingOne(3))

	// All in one func
	doubleAddOne := makeDoubler(addOne)
	fmt.Println("All concepts combined ::: ", doubleAddOne(1))
}

func addOne(a int) int {
	return a + 1
}
func addTwo(a int) int {
	return a + 2
}

// Function as an argumaent of another function
func printOperation(a int, f func(int) int) {
	fmt.Println("Function as argument of another :::", f(a))
}

// Function returns another function
func makeAdder(b int) func(int) int {
	return func(a int) int {
		return a + b
	}
}

// All the above concepts in one functions
func makeDoubler(f func(int) int) func(int) int {
	return func(a int) int {
		b := f(a)
		return b * 2
	}
}
