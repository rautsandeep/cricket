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
	var mu sync.Mutex
	for i := 0; i < 100; i++ {

	fmt.Println("Number of goroutine: ", runtime.NumGoroutine())
		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
//		fmt.Println("counter:", counter)
	}

	wg.Wait()

	fmt.Println("Number of goroutine: ", runtime.NumGoroutine())
	fmt.Println("counter:", counter)

}
