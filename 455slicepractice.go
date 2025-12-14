package main

import (
	"fmt"
)

func main() {

	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)
	fmt.Printf("The length of x is %v\n", len(x))
	fmt.Printf("the capacity of x is %v\n", cap(x))

	fmt.Println("appending valu to the slice x")

	x = append(x, 52)
	fmt.Println(x)

	x = append(x, 53, 54, 55)

	fmt.Println("slice x after appending 53,54,55")
	y := []int{56, 57, 58, 59, 60}
	fmt.Println(x)
	fmt.Println("appending the value from slice y to x")

	x = append(x, y[0:]...)
	fmt.Println("value of x post appending y ", x)

}
