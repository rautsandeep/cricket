package main

import (
	"fmt"
)

func conca(name string, s ...string) {

	fmt.Println(name, " friends name is", s)

}

func main() {

	fmt.Println("The name of sandeep friends will be given below")

	friend := []string{"sandeep", "Akshay", "Rahul", "Sagar", "Pramod", "swapnil"}
	conca("sandeep", friend...)

}
