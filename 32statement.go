package main

import (
	"fmt"
	"math/rand"
)

func main() {
	count := 0
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Println("The value of x is 3")
			count++
		}
	}
fmt.Println(count, " times the value of x set to 3")
}
