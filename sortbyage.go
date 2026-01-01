package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}
type byage []person

func (a byage) Len() int {
	return len(a)
}

// Swap is part of sort.Interface.
func (a byage) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (a byage) Less(i, j int) bool {
	return a[i].Age < a[j].Age
}
func main() {

	p1 := person{
		Name: "sandeep",
		Age:  22,
	}
	p2 := person{

		Name: "rahul",
		Age:  19,
	}
	p3 := person{
		Name: "shubham",
		Age:  21,
	}

	people := []person{p1, p2, p3}

	fmt.Println(people)
	sort.Sort(byage(people))
	fmt.Println(people)
}
