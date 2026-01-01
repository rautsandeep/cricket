package main

import (
	"fmt"
)

type mynumbers interface {
	int | float64
}

func main() {

	fmt.Printf("The sum of 3 and 5 is %v\n", addI(3, 5))
	fmt.Printf("The sum of 3.1 and 5.4 is %v\n", addF(3.1, 5.4))

	fmt.Printf("The sum of 3 and 5 is %v\n", generic(3, 5))
	fmt.Printf("The sum of 3.1 and 5.4 is %v and tyep %T\n", generic(3.1, 5.4), generic(3.1, 5.4))

	fmt.Printf("The sum of 3.1 and 5 is %v\n", generic(3.1, 5))
}

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {

	return a + b
}

func generic[T mynumbers](a, b T)T {
	return a + b
}
