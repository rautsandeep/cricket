package main

import (
	"fmt"
)

type person struct {
	firstname string
	lastname  string
	icecreams []string
}

func main() {
	p1 := person{
		firstname: "sandeep",
		lastname:  "snehal",
		icecreams: []string{"guava", "mango"},
	}

	p2 := person{

		firstname: "snehal",
		lastname:  "shinde",
		icecreams: []string{"strawberry", "santra"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Printf("%#v, %v\n", p1, p1)
	fmt.Printf("%#v, %v\n", p2, p2)
	for _, v := range p1.icecreams {
		fmt.Println(v)
	}

	for _, v := range p2.icecreams {
		fmt.Println(v)

	}

	m := map[string]person{

		p1.lastname: p1,
		p2.lastname: p2,
	}
	for k, v := range m {
		for _, j := range v.icecreams {
			fmt.Println(k, k, j)
		}
	}
}
