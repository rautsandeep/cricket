package main

import (
	"fmt"
)

func main() {

	c := make(chan []int, 1)
	xi := []int{42, 45}
	c <- xi

	fmt.Println(<-c)

	d := make(chan<- int, 2)
	d <- 40
	d <- 41
	/*
		fmt.Println(<- d)
		fmt.Println(<- d)
		fmt.Println(<- d)
	*/
	e := make(chan int, 4)
	e <- 25
	e <- 26
	e <- 27
	e <- 28
	close(e)
	for v := range e {
		fmt.Println(v)
	}
}
