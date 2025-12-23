package main

import (
	"fmt"
)

func foo() {

	fmt.Println("I am from foo function")
}

func bar() {

	fmt.Println("I am from the bar function")

}

func main() {
defer	foo()

	bar()

}
