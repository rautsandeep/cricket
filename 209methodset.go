package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

func (p *person) speak() {
	fmt.Printf("name is %v and age is %v\n", p.Name, p.Age)
}

type human interface {
	speak()
}

func saysomething(h human) {
	h.speak()
}

func main() {

	p1 := &person{
		Name: "sandeep",
		Age:  23,
	}

	p2 := person{
		Name: "snehal",
		Age:  22,
	}

	saysomething(p1)
	saysomething(&p2)
}
