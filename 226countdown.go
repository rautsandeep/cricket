package main

import (
	"fmt"
	"sync"
)

func main() {
	a := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 100; i++ {
			a <- i
		}
		close(a)
		wg.Done()
	}()

	go func() {
		for v := range a {

			fmt.Println(v)
		}
		wg.Done()

	}()
	wg.Wait()
}
