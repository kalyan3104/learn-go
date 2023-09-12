package main

import (
	"fmt"

	"github.com/RashikAnsar/learn-go/7.Packages/3.third_party_packages/mapper"
)

func main() {
	fmt.Println(mapper.Greet("Howdy, what's new?"))
	fmt.Println(mapper.Greet("Comment allez vous?"))
	fmt.Println(mapper.Greet("Wie geht es Ihnen?"))
}
