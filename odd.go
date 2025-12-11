package main

import (
	"fmt"
)

func main() {

	fmt.Println("Trying to pring odd numbers between the 1 to 100")
	count:=0
	for i := 1; i < 101; i++ {
		remainder := i % 2
		if remainder != 0 {
			fmt.Println(i)
			count++
		}

	}
	fmt.Println("Total count of the odd number between 1 to 100 is", count)
}
