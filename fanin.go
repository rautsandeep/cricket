package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

go 	send(even, odd)
go	receiver(even, odd, fanin)

	for v1 := range fanin {
		fmt.Println(v1)
	}

	fmt.Println("program about to exit")
}

func send(even, odd chan<- int) {

	for i := 0; i < 10; i++ {

		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}

	}
		close(even)
		close(odd)
}

func receiver(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v1 := range odd {
			fanin <- v1
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
