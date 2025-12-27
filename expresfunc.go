package main

import (
	"fmt"
)

func addition(s, j int) int {
	fmt.Println("This is named funnction")
	return s + j
}

func main() {
	func(s, j int) int {
		fmt.Println("this is anonymous subtraction function")
		return s - j

	}(10, 4)

	y := func(a, b int) int {
		return a * b
	}

	fmt.Println("the multiplication of two numbres is", y(3, 4))
}
