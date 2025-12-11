package main

import (
	"fmt"
)

func main() {
	var y int = 100
	for y > 0 {
		fmt.Println(y)
		y -= 1
		if y < 50 {
			break
		}
	}
}
