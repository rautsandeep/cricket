package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Print("Provide the word as input: ")
	fmt.Scanln(&input)
	slice := []byte(input)
	x := Ispalindrome(slice)
	if x {
		fmt.Printf("%v is palindrome\n", input)
	} else {

		fmt.Printf("%v is not  palindrome\n", input)
	}
}

// Ispalindrome function checks if the string is palindrome and returns boolean value of true or false.
func Ispalindrome(a []byte) bool {
	b := make([]byte, 0)
	count := len(a) - 1
	for i := count; i >= 0; i-- {
		b = append(b, a[i])

	}
	fmt.Println(a)
	fmt.Println(b)

	if string(a) == string(b) {
		return true
	} else {
		return false
	}

}
