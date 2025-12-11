package main

import (
	"fmt"
)

func main() {

	xi := []int{1, 11, 21, 31}
	for i, v := range xi {
		fmt.Printf("Value is %v for index %v\n", v, i)
	}
	m := map[string]int{
		"sandeep": 1,
		"snehal":  2,
	}
	for key, value := range m {
		fmt.Printf("The value %v for key is %v\n", value, key)
	}
	for key := range m {
		fmt.Printf("the list of keys are %v\n", key)
	}
	for _,value := range m {
		fmt.Printf("the list of keys are %v\n", value)
	}
}
