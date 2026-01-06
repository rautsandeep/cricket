package main

import (
	"fmt"
)

func main() {

	fmt.Println("We will check the directional channel")

	a := make(chan int, 4)
	b := make(<-chan int, 3)
	c := make(chan<- int, 4)

	a <- 40
	a<-50

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	b = a

fmt.Println(<-a)
fmt.Println(<-b)
fmt.Printf("%T\n",(<-chan int)(a) )
	fmt.Printf("%T\n", b)
}
