package main

import (
	"fmt"
)

type person struct {
	first string
	age   int
}

func (p person) speak() {
	fmt.Printf("The name of the person is %v and age is %v\n", p.first, p.age)
}
func main() {

	p1 := person{
		first: "sandeep",
		age:   33,
	}

	p2 := person{
		first: "snehal",
		age:   32,
	}

	p1.speak()
	p2.speak()

}
