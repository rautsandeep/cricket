package main

import (
	"fmt"
)

func main() {

	a := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Printf("slice a %#v\n", a)

	for _, v := range a {

		fmt.Printf("the element stored in the slice a at each index position is %v of the type %T\n", v, v)
	}

		fmt.Printf("%T %#v\n",a,a)
		fmt.Printf("%T %v\n",a,a)

}
