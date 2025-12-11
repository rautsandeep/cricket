package main

import (
	"fmt"
	"math/rand"
)

func main() {
	for i := 0; i < 40; i++ {
		x := rand.Intn(3)
		fmt.Printf("The name of the variable is x and value is %v\n", x)

		if x <= 100 {
			fmt.Println("Between 0 and 100")
		} else if x > 100 && x <= 200 {

			fmt.Println("between 101 and 200")
		} else if x > 200 && x <= 250 {

			fmt.Println("between 201 and 250")
		}
	}

}
