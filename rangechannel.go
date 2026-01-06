package main

import (
	"fmt"
)

func main() {
	fmt.Println("channel ranging")

	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {

			c <- i
		}
//		close(c)
	}()

	for v := range c {
		fmt.Println("iteration number: ", v)
	}

	fmt.Println("Program exited")

}
