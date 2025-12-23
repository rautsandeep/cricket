package main

import (
	"fmt"
)

func sum(s ...int) int {
	fmt.Println(s)

	fmt.Printf("the type of %v is %T\n", s, s)

	var sum int
	for _, v := range s {
		sum += v
	}
return sum
}
func main() {
	fmt.Println("We will calculate the sum of all numbers providec")

	y := sum(23, 34, 593, 29, 3)
	fmt.Println("The sum of all numbers is ", y)


	z:= sum()
	fmt.Println(z)
}

