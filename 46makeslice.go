package main

import (
	"fmt"
)

func main() {

	xi := []int{2, 3, 4}
	fmt.Println(xi)
	fmt.Printf("%#v\n", xi)
	fmt.Println("The length of slice is ", len(xi))
	fmt.Println("The capacity of slice is ", cap(xi))

	xs := make([]int, 0, 10)
	fmt.Println(xs)
	fmt.Printf("%#v\n", xs)
	fmt.Println("The length of slice is ", len(xs))
	fmt.Println("The capacity of slice is ", cap(xs))

	xs = append (xs, 4,5,4,5,5,6,2,6,3,7,8,4,4,3,5,6,5,0,6,7)
	fmt.Printf("%#v\n",xs)
	fmt.Println("length after appednding few more elements ", len(xs))
	fmt.Println("capacity after appending few more elements", cap(xs))


	xs = append (xs, 5,0,6,7)
	fmt.Printf("%#v\n",xs)
	fmt.Println("length after appednding few more elements ", len(xs))
	fmt.Println("capacity after appending few more elements", cap(xs))

	xa:=make([]int,5)
	fmt.Println(xa)
	fmt.Println(len(xa))
	fmt.Println(cap(xa))

	xa= append(xa, 3,3)
	fmt.Println(xa)
	fmt.Println(len(xa))
	fmt.Println(cap(xa))
}
