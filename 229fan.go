package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	c := make(chan int)
	gs := 10
	go func() {

		wg.Add(gs)
		for i := 1; i <= gs; i++ {

			go func() {
				m:= i*10
				for j := m; j <= m+10; j++ {
					c <- j
				}
				wg.Done()
			}()
		}
		wg.Wait()
		close(c)
	}()
	for v := range c {
		fmt.Println(v)

	}
}
