package main

import (
	"fmt"
)

func main() {

	xi := []int{42, 43, 44, 45, 46, 47}

	for i, value := range xi {
		fmt.Printf("The value belongs to index %v is %v\n", i, value)
	}

}
