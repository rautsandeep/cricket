package main

import (
	"fmt"
)

func main() {
	x := 10
	y := &x
	fmt.Println(y)
	fmt.Println(&y)
	fmt.Println(*y)

}
