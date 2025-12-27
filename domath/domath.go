package main

import (
	"fmt"
)

func main() {

	z := domath(40, 5, add)
	fmt.Println("Action performed on the operands resulted in ", z)
}

func domath(a int, b int, f func(x int, y int) int) int {
	return f(a, b)
}
func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}
