package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	in := make(chan int)
	out := make(chan int)

	go fanin(in)
	go fanoutin(in, out)
	for v := range out {
		fmt.Println(v)

	}
}

func fanin(in chan<- int) {

	for i := 0; i < 100; i++ {
		in <- i
	}
	close(in)
}

func fanoutin(in <-chan int, out chan<- int) {

	var wg sync.WaitGroup

	for v := range in {
		wg.Add(1)
		go func(v2 int) {
			out <- work(v2)
			wg.Done()
		}(v)
		wg.Wait()
	}
	close(out)
}

func work(v2 int) int {
	time.Sleep(3 * time.Second)
	return v2 * 5

}
