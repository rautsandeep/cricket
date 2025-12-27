package main

import (
	"fmt"
)

func main() {
	y := outer()
	y(8)
	y(3)
	y(5)
	y(7)
	y(3)
	y(5)

}

func outer() func(a int) int {
	fmt.Println("This is outer function")
	s := 0
	return func(a int) int {
		s +=a
		fmt.Println("the value of s is", s)
		return s
	}
}
