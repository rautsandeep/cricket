package main

import (
	"fmt"
)

func main() {

	var a [3]int
	fmt.Println(a)
	a[1] = 5
	a[0] = 10

	fmt.Println(len(a))
	fmt.Println(a)

	b := [3]string{"sandeep", "Snehal"}
	fmt.Println(b)
	x := len(b)
	y := 0
	for i := 0; i < x; i++ {
		if b[i] == "Snehal" {

			fmt.Println("The index of the position where snehal resdes is", i)
			y = 1
			break
		}
	}
		if y == 0 {
			fmt.Println("The literal snehal does not exists in the array with name b")
		}

	fmt.Println("the length of array b is ", len(b))
	var c [3]string
	c = b
	fmt.Println("lenth of arrat c is", len(c))
	fmt.Println(c)

	d := [...]int{5, 5, 5, 5}
	fmt.Println("The length of the d is ", len(d))

	fmt.Printf("The type of array is %T\n", a)
	fmt.Printf("The type of array is %T\n", b)
	fmt.Printf("The type of array is %T\n", c)
	fmt.Printf("The type of array is %T\n", d)
}
