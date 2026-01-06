package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := foo(q)

	bar(c, q)
	fmt.Println("about to exit")

}

func foo(q chan<- int) <-chan int {

	c := make(chan int)
	go func() {

		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		q <- 1
	}()

	return c

}

func bar(c <-chan int, q <-chan int) {

	for {

		select {

		case v := <-c:
			fmt.Println(v)
		case v := <-q:
			fmt.Println(v)
			return
		}
	}
}
