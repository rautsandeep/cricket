package main

import (
	"fmt"
)

func main() {
	a := "sandeep"
	b := []byte(a)
	fmt.Println(a)
	fmt.Println(b)
	c := []byte{115, 97, 110, 100, 101, 101, 112}
	fmt.Println(string(c))

}
