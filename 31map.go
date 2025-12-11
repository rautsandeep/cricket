package main

import (
	"fmt"
)

func main() {
	age := map[string]int{
		"sandeep": 33,
		"snehal":  32,
		"foo":     15,
		"bar":     45,
	}
	subage := age["snehal"]
	fmt.Println(subage)
	var i int
	for _, value := range age {
		if value < 18 {
			i++
		}

	}
	fmt.Println("the count of underage voter is", i)
	if v, ok := age["bar"]; ok {
		fmt.Println(ok)
		fmt.Println(v)
	}
	if v, ok := age["gar"]; !ok {
		fmt.Println(ok)
		fmt.Println(v)
	}
}
