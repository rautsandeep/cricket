package main

import (
	"fmt"
)

type person struct {
	first  string
	last   string
	age    int
	gender string
}
type student struct {
	person
	class int
}

func main() {

	p1 := person{
		first:  "sandeep",
		last:   "raut",
		age:    20,
		gender: "male",
	}

	p2 := person{
		first:  "snehal",
		last:   "shinde",
		age:    20,
		gender: "female",
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.first, p2.last)
	hushar := student{
		person: person{
			first:  "omkar",
			last:   "tetame",
			age:    15,
			gender: "male",
		},
		class: 9,
	}
	fmt.Println(hushar)
	fmt.Println(hushar.class, hushar.person.first, hushar.person.last)
	fmt.Printf("%#V %T \n", hushar)
}
