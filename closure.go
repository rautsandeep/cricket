package main

import (
	"fmt"
)

func main() {
	y := incrementor()
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())

}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
