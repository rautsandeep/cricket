package main

import (
	"fmt"
)

func main() {

	for i := 0; i < 100; i++ {

		y := i % 2

		if y != 0 {
			if i == 51 {
				fmt.Println("the value of i is 51. Need to skip this iteration")
				continue
			}
			fmt.Printf("%v is odd number\n", i)

		}
	}
}
