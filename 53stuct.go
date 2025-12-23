package main

import (
	"fmt"
)

func main() {

	type person struct {
		firstname string
		lastname  string
		icecream  []string
	}

	p1 := person{
		firstname: "sandeep",
		lastname:  "raut",
		icecream:  []string{"vanila", "blueberry"},
	}
	p2 := person{

		firstname: "snehal",
		lastname:  "shinde",
		icecream: []string{"strawberry", "chocolate"},
	}

	fmt.Println(p1.firstname, p1.lastname, p1.icecream)
	fmt.Println(p2.firstname, p2.lastname, p2.icecream)
	for _,v := range p1.icecream {
		fmt.Println(v)
	}

	for _,v := range p2.icecream {
		fmt.Println(p1.firstname,"'s favourite icecreams are",v)
	}
	fmt.Printf("%#v\n", p1)
	fmt.Printf("%#v\n", p2)

}
