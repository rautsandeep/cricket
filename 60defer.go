package main

import (
	"fmt"
)

func main() {
defer 	foo()
	 bar()
	defer car()
	foo()

}

func foo() {
	fmt.Println("The output is from the foo function")

}

func bar() {

	fmt.Println("This is output from the bar function")
}

func car() {

	fmt.Println("This is output from the ccccaaaarrrrr function")
}
