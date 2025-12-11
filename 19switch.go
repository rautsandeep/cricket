package main

import (
	"fmt"
	"math/rand"
)

func main() {

	z := rand.Intn(250)
	switch {
	case z <= 100:
		fmt.Printf("The value of z is %v which is between 0 and 100\n", z)
	case z > 100 && z <= 200:
		fmt.Printf("The value of z is %v which is between 101 and 200\n", z)
	default:
		fmt.Printf("The value of z is %v which is between the 201 and 250\n", z)
	}

}
