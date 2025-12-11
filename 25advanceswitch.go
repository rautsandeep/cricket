package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 42; i++ {
		x := rand.Intn(5)
		fmt.Printf("Iteration number:%v The value of x is %v\n",i, x)

		switch x {

		case 0:
			fmt.Println("the value of x is 0")
		case 1:
			fmt.Println("the value of x is 1")
		case 2:
			fmt.Println("The value of x is 2")
		case 3:
			fmt.Println("The value of x is 3")
		default:
			fmt.Println("The value of x is 4")
		}
	}
}
