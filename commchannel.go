package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int, 3)
	wg.Add(2)
	go foo(c)

	go bar(c)
	wg.Wait()
	fmt.Println("about to exit")
}

func foo(c chan<- int) {

	c <- 45
	wg.Done()
}

func bar(c <-chan int) {
	fmt.Println(<-c)
	wg.Done()
}
