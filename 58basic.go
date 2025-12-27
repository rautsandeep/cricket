package main

import (
	"fmt"
)

func main() {

	p := foo(23, 33)
	q, name := bar(50, "sandeep")
	fmt.Println("sum  of the number are", p)
	fmt.Printf("Age of the %v after 7 decades is %v\n", name, q)
}
func foo(a, b int) int {

	return a + b
}

func bar(b int, s string) (int, string) {

	fmt.Printf("The age of the %v is %v\n", s, b)
	b = b * 7
	return b, s
}
