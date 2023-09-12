package main

import "fmt"

func main() {
	s := "Hello, world"
	// s2 := s[0:5]
	// s3 := s[7:12]
	s4 := s[:5]
	s5 := s[7:]
	fmt.Println(s, len(s), s4, len(s4), s5, len(s5))
}
