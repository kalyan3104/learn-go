package main

import "fmt"

// Foo contains int and string
// int is a public field
type Foo struct {
	A int
	b string
}

func main() {
	f := Foo{}
	fmt.Println(f)

	g := Foo{10, "Hello"}
	fmt.Println(g)

	h := Foo{
		b: "Goodbye",
	}
	fmt.Println(h)

	h.A = 1000
	fmt.Println(h.A)

	fmt.Println("=========================")
	f1 := Foo{
		A: 20,
	}
	var f2 Foo
	f2 = f1
	f2.A = 100
	fmt.Println(f2.A)
	fmt.Println(f1.A)

	var f3 *Foo = &f1
	f3.A = 200
	fmt.Println(f3.A)
	fmt.Println(f1.A)
}
