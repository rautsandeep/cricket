package main

import (
	"fmt"
)

func main() {

	name := "sandeep"
	surname := "Raut"

	n := &name
	p := &surname

	fmt.Printf("name and surname is %v and %v \n", n, p)
	fmt.Printf("name and surname is %T and %T \n", n, p)

}
