// A function that is the part of the definition of a data-type
// Go doesnot have inheritance

package main

import "fmt"

// Foo has a int and string public fields
type Foo struct {
	A int
	B string
}

// Print the string
func (f Foo) String() string {
	return fmt.Sprintf("A: %d and B: %s", f.A, f.B)
}

// Double the value
func (f *Foo) Double() {
	f.A = f.A * 2
}

func main() {
	f := Foo{
		A: 10,
		B: "Hello",
	}
	fmt.Println(f.String())
	f.Double()
	fmt.Println(f.String())
}
