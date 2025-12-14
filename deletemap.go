package main

import (
	"fmt"
)

func main() {

	colour := map[string]string{
		"sandeep": "blue",
		"snehal":  "pink",
		"omkar":   "white",
	}
	fmt.Println("who likes which colour\n", colour)

	fmt.Println("Adding the colour liked by the Priyank")

	colour["priyank"] = "green"
	fmt.Println(colour)

	fmt.Println("deleting the key value from the map")

	delete(colour, "omkar")
	fmt.Println(colour)
	delete(colour, "omkar")
	fmt.Println("The value of the colour liked by omkear is ",colour["omkar"])
	fmt.Println(colour["omkar"])
}

