// Empty interface

package main

import "fmt"

func main() {
	var i interface{}
	i = "Hello"
	j := i.(string)
	k, ok := i.(int)
	fmt.Println(j, k, ok)
	m := i.(int) //	panic: interface conversion: interface {} is string, not int
	fmt.Println(m)
}
