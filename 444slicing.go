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

	fmt.Printf("%T %#v\n", a, a)
	fmt.Printf("%T %v\n", a, a)

	x := []int{}
	x = a[:5]

	fmt.Println("The value of the x is ", x)
	y := make([]int, 0)
	y = a[5:]
	fmt.Println(y)

	z := make([]int, 0, 5)
	z = a[2:7]

	fmt.Println(z)
	p := make([]int, 0)
	p = a[1:6]
	fmt.Println(p)

}
