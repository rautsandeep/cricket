package main

import (
	"fmt"
)

func main() {

	p := []int{2, 5, 488, 594, 59, 599, 3049, 485}

	sum := foo(p...)
	fmt.Println("The sum of all the elements of slice is ", sum)

	barsum := bar(p)
	fmt.Println("The sum obtained by bar function is ", barsum)

}

func foo(x ...int) int {
	sum := 0
	for _,v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {

	sum := 0
	for _,v := range x {

		sum += v
	}
	return sum
}
