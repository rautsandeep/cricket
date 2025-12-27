package main

import (
	"fmt"
)

func main() {
	x := add(5, 34)
	fmt.Println("the sum of two number is", x)
}

func add(a, b int) int {

	return a + b
}
