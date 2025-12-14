package main

import (
	"fmt"
)

func main() {

	var a [5]int

	fmt.Println("the value of each element in a ", a)

	a[0] = 23
	a[1] = 24
	a[2] = 25
	a[3] = 26
	a[4] = 27

	fmt.Printf("%#v\n", a)

	for i, v := range a {

		fmt.Printf("The value at index position %v is %v\n", i, v)

	}

}
