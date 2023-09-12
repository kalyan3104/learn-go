package main

import "fmt"

//Foo has a public int and private string
type Foo struct {
	A int
	b string
}

// Bar has a public string and has embedded Foo
type Bar struct {
	C string
	F Foo
}

// Baz has a public string and has embedded Foo
type Baz struct {
	D string
	Foo
}

func main() {
	f := Foo{A: 10, b: "Hello"}
	b1 := Bar{C: "Fred", F: f}
	fmt.Println(b1.F.A)
	b2 := Baz{D: "Nancy", Foo: f}
	fmt.Println(b2.A)

	var f2 Foo = b2.Foo
	fmt.Println(f2)
}
