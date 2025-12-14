package main

import (
	"fmt"
)

func main() {
	a := []string{"james", "bond", "shaken not stirred"}
	fmt.Println(a)
	b := []string{"miss", "Moneypenny", "I am 008"}

	fmt.Println(b)
	c := [][]string{a, b}

	for i := 0; i < len(c); i++ {
		for j := 0; j < len(c[i]); j++ {
			fmt.Println(c[i][j])
		}
	}

}
