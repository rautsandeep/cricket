package main

import (
	"fmt"
)

func foo() {
	fmt.Println("hi the statement has been called from the Foo function")
}

func bar(a string) {

	fmt.Printf("The award goes to %v\n", a)
}
func summul(a, b int) (x, y int) {
	sum := a + b
	mul := a * b
	return sum, mul
}
func votecount(s string, count int) (string, int) {
	count *= 7
	return fmt.Sprint("total count of voters in ", s,  "assembly is"), count
}

func main() {
	fmt.Println("the execution of main program started")
	foo()
	bar("Snehal")
	fmt.Println("the end of func program")
	p, q := summul(5, 3)
	fmt.Printf("Sum of two numbers are %v\n", p)
	fmt.Printf("Multiplication of two num are %v\n", q)
	m, n := votecount("satara", 30)
	fmt.Println(m, n)
}
