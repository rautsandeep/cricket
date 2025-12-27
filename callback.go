package main

import (
	"fmt"
)

func main() {

	x := sum(20, 5)
	fmt.Println("the sum of two numbbers is ", x)

	y := substract(25, 5)
	fmt.Println("the substraction of two number is", y)

	z := doMath(4, 5, sum)
	fmt.Println("output of doMath function is ", z)

	fmt.Printf("%T\n", sum)
	fmt.Printf("%T\n", substract)
	fmt.Printf("%T\n", doMath)

}

func sum(a int, b int) int {
	return a + b
}

func substract(a int, b int) int {

	return a - b
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}
