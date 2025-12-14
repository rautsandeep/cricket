package main

import (
	"fmt"
)

func main() {

	books := map[string][]string{

		"bond_james":       []string{"shaken, not stirred", "martinis", "fast cars"},
		"moneypenny_jenny": []string{"intelligence", "literature", "computer science"},
		"no_dr":            []string{"cats", "ice cream", "sunsets"},
	}
	fmt.Println(books)
	books["fleming_ian"]= []string{"steaks", "cigars", "espionage"}
	fmt.Println(books)
	for k, v := range books {

		fmt.Printf("Printing the value for %v\n", k)
		for x, y := range v {
			fmt.Printf("%v %v\n", x, y)
		}
	}

	fmt.Println("--------------------------------------")
	delete(books, "moneypenny_jenny")
	fmt.Println(books)
}
