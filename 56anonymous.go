package main

import (
	"fmt"
)

func main() {

	p2 := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "sandeep",
		friends: map[string]int{
			"Avinash": 33,
			"Ajay":    35,
			"Rohit":   34,
		},
		favDrinks: []string{"coke", "pepsi"},
	}

	for k, v := range p2.friends {
		fmt.Println(p2.first, "is friend of ",k, v)
	}
	for _,v:= range p2.favDrinks {
fmt.Println(p2.first,"likes drink", v)
	}

}
