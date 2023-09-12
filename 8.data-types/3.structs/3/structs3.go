package main

import (
	"encoding/json"
	"fmt"
)

// Person has a string(Name) and int(Age) field
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	rashik := `{ "name": "Rashik", "age": 20}`
	var b Person
	json.Unmarshal([]byte(rashik), &b)
	fmt.Println(b)
	rashik2, _ := json.Marshal(b)
	fmt.Println(string(rashik2))
}
