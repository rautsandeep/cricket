package main

import (
	"fmt"
)

func main() {
	func(i int) int {
		fmt.Println(i)
		return i*5
	}(5)

}
