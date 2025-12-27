package main

import (
	"fmt"
)

func main() {

	p := snehal(23, 4, add)
	fmt.Println("addition of two ", p)
	q := snehal(23, 4, sub)
	fmt.Println("subtraction of two number is ", q)
	r := snehal(23, 4, mul)
	fmt.Println("Multiplication of two numbers:", r)

	t := snehal(23, 4, mod)
	fmt.Println("MOdulas of ", t)
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func mul(a, b int) int {
	return a * b
}

func mod(a, b int) int {
	return a % b
}

func snehal(x int, y int, f func(int, int) int) int {
	return f(x, y)
}
