package main

import (
	"fmt"
)

func main() {
	x := 44
	fmt.Println(x)
	fmt.Println(&x)
	p := &x
	fmt.Println(*p)


	fmt.Printf("%T\n",&x)
}
