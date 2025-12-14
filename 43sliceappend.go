package main

import (
	"fmt"
)

func main() {

	xi := []int{4, 5, 5, 7}
	fmt.Println("before appending item to slice xi", xi)

	xi = append(xi, 8, 9, 10)
	fmt.Println("after apending values to the slice", xi)
	fmt.Printf("%#v\n",xi)

}
