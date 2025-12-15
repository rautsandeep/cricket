package main

import (
	"fmt"
)

func main() {

	p := struct {
		animal   string
		bird     string
		count    string
		teamsize int
	}{
		animal:   "dog",
		bird:     "parrot",
		count:    "six",
		teamsize: 5,
	}

	fmt.Println(p)
	fmt.Printf("%v, %#v\n", p,p)

}
