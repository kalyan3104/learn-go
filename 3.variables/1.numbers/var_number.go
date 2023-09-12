package main

import "fmt"

func main() {
	var i int8 = 8
	var j int16 = 16
	var k int = 3264 // either int32 or int64 based on CPU architecture
	var q rune = 8  // It can also be reffered as int32

	var l uint8 = 8 // It can also be reffered as byte 
	var m uint16 = 16
	var n uint = 3264

	var o float32 = 32.32
	var p float64 = 64.64

	fmt.Println("int type values: ", i, j, k)
	fmt.Println("uint type values: ", l, m, n)
	fmt.Println("int32 and rune are same: ", q)
	fmt.Println("float type values: ", o, p)

	// If we want to perform any mathematical operator it must bbe of same type
}
