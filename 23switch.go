package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x, y := rand.Intn(10), rand.Intn(10)

	fmt.Printf("The value of x is %v and the value of y is %v\n", x, y)

	switch {
	case x < 4 && y < 4:
		fmt.Printf("Both x and y are less than 4\n")
	case x > 6 && y > 6:
		fmt.Printf("x and y are %v and %v respectively and both are greater tha 6\n", x, y)
	case x >= 4 && x <= 6:
		fmt.Printf("the value of x is %v and greater than 4 and less than6\n", x)
	case y != 5:
		fmt.Printf("The vaue of the y is %v which is other than 5\n", y)
	default:

		fmt.Println("None of the previous cases were met")

	}

}
