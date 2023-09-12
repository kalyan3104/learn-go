package main

import "fmt"

func main() {
	fmt.Println("==========================")
	m := make(map[string]int)
	m["hello"] = 300
	h := m["hello"]
	fmt.Println("hello in m:", h)
	fmt.Println("a in m:", m["a"])

	if v, ok := m["hello"]; ok {
		fmt.Println("hello in m:", v)
	}

	m["hello"] = 20
	fmt.Println("hello in m:", m["hello"])

	fmt.Println("==========================")
	m2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 50,
	}

	for k, v := range m2 {
		fmt.Println(k, v)
	}

	fmt.Println("b in m2:", m2["b"])
	delete(m2, "b")
	fmt.Println("b in m2:", m2["b"])

	fmt.Println("==========================")
	m4 := map[string]int{
		"a": 1,
		"b": 2,
	}
	var m3 map[string]int

	fmt.Println("goodbye in m4:", m4["goodbye"])
	m3 = m4
	m3["goodbye"] = 400
	fmt.Println("goodbye in m3:", m3["goodbye"])
	fmt.Println("goodbye in m4:", m4["goodbye"])

	modMap(m4)
	fmt.Println("cheese in m4:", m4["cheese"])
}

func modMap(m map[string]int) {
	m["cheese"] = 20
}
