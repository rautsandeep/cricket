package main

import (
	"fmt"
)

func main() {

	x := func(a, b int) int {
		fmt.Println("Value of a is", a)
		fmt.Println("Value of b is", b)
		return a + b

	}

	y:=	x(4, 9)
	fmt.Println("sum is",y)

}
