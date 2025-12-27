package main

import (
	"fmt"
)

func main() {

	fmt.Println("testing the dereferencing with add function")
	p := 5
	q := 40
	y := add(&p, &q)
	fmt.Println("the value of sum is", *y)
	xi := []int{2, 4, 3, 5, 8}

	fmt.Println(xi)
	slice(xi)
	fmt.Println(xi)
}

func add(a *int, b *int) *int {
	first := *a
	second := *b
	sum := first + second
	return &sum
}

func slice(ii []int) {
	ii[0] = 5
}
