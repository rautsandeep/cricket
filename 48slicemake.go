package main

import (
	"fmt"
)

func main() {

	a := []int{4, 5, 6, 2, 0, 8}

	fmt.Println("the value of array a is ", a)

	b := make([]int, 6)
	fmt.Println(b)

	fmt.Println("the value of array a is ", a)
	fmt.Println("Now copying the content of a to b")

	    copy(b, a)
	   fmt.Println("after copying b from a ", b)

	   a[0] = 9
	   fmt.Println(a)
	   fmt.Println(b)
}
