package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("CPU Count: ", runtime.NumCPU())
	fmt.Println("Number of goroutine: ", runtime.NumGoroutine())
	counter := 0
	gs := 100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < 100; i++ {

	fmt.Println("Number of goroutine: ", runtime.NumGoroutine())
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
//		fmt.Println("counter:", counter)
	}

	wg.Wait()

	fmt.Println("Number of goroutine: ", runtime.NumGoroutine())
	fmt.Println("counter:", counter)

}
