package main

import (
	"fmt"
)

func sum(s ...int) int {
	addition := 0
	for _, v := range s {
		addition += v

	}
	return addition

}
func main() {

	x := []int{40, 35, 33, 40, 40, 10}

	sum := sum(x...)
	fmt.Println("The sum of all the elements of the slice is", sum)

}
