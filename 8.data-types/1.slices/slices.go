package main

import "fmt"

//A slice is a growable sequence of values of a single specified type
func main() {
	fmt.Println("===================================")
	s := make([]string, 0)
	fmt.Println("length of s:", len(s))
	s = append(s, "hello")
	fmt.Println("length of s:", len(s))
	fmt.Println("contents of s[0]:", s[0])
	s[0] = "goodbye"
	fmt.Println("contents of s[0]:", s[0])

	fmt.Println("===================================")
	s2 := make([]string, 2)
	fmt.Println("contents of s2[0]:", s2[0])
	s2 = append(s2, "hello")
	fmt.Println("contents of s2[0]:", s2[0])
	fmt.Println("contents of s2[2]:", s2[2])
	fmt.Println("length of s2:", len(s2))

	fmt.Println("===================================")
	s3 := []string{"a", "b", "c"}
	for k, v := range s3 {
		fmt.Println(k, v)
	}

	fmt.Println("===================================")
	s4 := s3[0:2]
	fmt.Println("s4: ", s4)
	s3[0] = "d"
	fmt.Println("s4: ", s4)

	fmt.Println("===================================")
	var s5 []string
	s5 = s3
	s5[1] = "Camel"
	fmt.Println("s5: ", s5)
	fmt.Println("s3: ", s3)

	modSlice(s3)
	fmt.Println("s3[0]: ", s3[0])
}

func modSlice(s []string) {
	s[0] = "hello"
}
