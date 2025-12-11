package main

import (
	"fmt"
	//"math"
)

func main() {
	for i := 2; i < 6; i++ {
		fmt.Println(i)
	}
	y := 5
	for {
		fmt.Println(y)
		if y == 40 {
			break
		}
		y++
	}
	standard := 1
	for {
		fmt.Println("I am in standard", standard)
		standard++
		if standard > 11 {
			break
		}
		if standard == 5 {
			fmt.Println("skipping 5th standard")
			continue
		}
	}
}
