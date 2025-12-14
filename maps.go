package main

import (
	"fmt"
)

func main() {

	age := map[string]int{
		"shubham": 17,
		"omkar":   15,
		"sandeep": 12,
		"snehal":  12,
	}
	fmt.Printf("the age of the shubham is %v\n", age["shubham"])

	for k, v := range age {
		fmt.Printf("The age of %v is %d\n", k, v)
	}

	fmt.Println("Value of the map is ",age)

	marks:= make(map[string]int)
	fmt.Println(marks)
	marks["shubham"]= 80
	marks["omkar"]= 65
	marks["snehal"]= 94
	marks["sandeep"]= 70
	fmt.Println( marks)


	fmt.Printf("lenght of the map is %v\n", len(marks))
	fmt.Println("printing only keys from map")

	for k := range marks {
fmt.Println(k)
	}

	for _,v := range marks {
fmt.Println(v)
	}
	
	friends:= []string{"Akshay", "swapnil", "Rahul", "sunil", "sagar", "Pramod"}
	fmt.Printf("Sandeep has %v number of friends\n", len(friends))
	for _, v:= range friends {
fmt.Println(v)
	}
}
