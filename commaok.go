package main

import (
	"fmt"
)

func main() {

	run := map[string]int{
		"sachin": 15000,
		"kohali": 14000,
		"dravid": 16000,
		"data":   12000,
	}
	fmt.Println(run)

	v, ok := run["surya"]
	if ok {
		fmt.Println("The run scored by surya", v)
	} else {

		fmt.Println("The key surya does not exists")
	}

	run["surya"]= 2000
	t, ok := run["surya"]
	if ok {
		fmt.Println("The run scored by surya", t)
	} else {

		fmt.Println("The key surya does not exists")
	}


}

