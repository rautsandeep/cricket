package main

import (
	"fmt"
)

func main() {

	xi := []int{5, 3, 5, 1, 3, 4, 9, 0, 4, 3, 9}
	fmt.Printf("%#v\n", xi)
	fmt.Println()
	fmt.Println("------------------")
	fmt.Println()
	fmt.Println("deleting all the entries of 3 from the slice")
	count := 0
	for i := range xi {
		if xi[i] == 3 {
			count++
		}
	}
	fmt.Println("The number of times 3 present in slice is ", count)

	for i := 0; i < count; i++ {
		for j, v := range xi {
			if v == 3 {

				k := j

				xi = append(xi[:k], xi[k+1:]...)
				fmt.Println(xi)

			}
		}

	}

}
