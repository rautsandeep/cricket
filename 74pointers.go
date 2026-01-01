package main

import (
	"fmt"
)

func main() {

	name := "sandeep"
	fmt.Println("My name is ", name)
	fmt.Println("memory address of the location where the value of name stored is", &name)
	addname := &name
	fmt.Println("value stored at the memory location of name is", *addname)
}
