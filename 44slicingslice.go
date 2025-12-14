package main

import (
	"fmt"
)

func main() {

	fmt.Println("declaring and initializing the slice")

	xi := []int{5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Println(xi)
	fmt.Printf("%#v\n", xi)
	fmt.Println()
	fmt.Println("------------------")

	fmt.Println()

	fmt.Println("slicing the elemements from 2 to 5")

	fmt.Printf("%#v\n", xi[2:6])

	fmt.Println()
	fmt.Println("----------------")
	fmt.Println()
	fmt.Printf("slicing the slice from the fist element to the 6")
	fmt.Printf("%#v\n", xi[:7])

	fmt.Println()
	fmt.Println("-----------------------")
	fmt.Println()
	fmt.Println("Slicing the slice from the 4th to last element")

	fmt.Printf("%#v\n", xi[4:])
}
