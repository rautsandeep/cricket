package main

import (
	"fmt"
	"sort"
)

func main() {

	xi := []int{21, 43, 5, 3, 59, 203, 44, 394}
	xs := []string{"sandeep", "snehal", "rahul", "shubham", "sachin", "dheeraj"}

	sort.Ints(xi)
	fmt.Println("The sorted string is ", xi)

	sort.Strings(xs)
	fmt.Println("The sorted string slice is ", xs)
}
