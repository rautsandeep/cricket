package main

import (
	"fmt"
)

func main() {

	a := []int{0, 4, 3, 4, 7}
	b := a
	fmt.Println("the value of slice a is ", a)
	fmt.Println("The value of slice b is ", b)

	a[0] = 9

	fmt.Println("the value of slice a is ", a)
	fmt.Println("The value of slice b is ", b)

}
