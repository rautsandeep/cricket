package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup

func main() {

	const gs = 100
	fmt.Println("gs:", gs)
	fmt.Println("Number of CPU: ", runtime.NumCPU())
	fmt.Println("GoRoutine: ", runtime.NumGoroutine())
	wg.Add(gs)
	var counter int64
	fmt.Println("counter:", counter)
	for i := 0; i < gs; i++ {

		go func() {
			atomic.AddInt64(&counter, 1)
			fmt.Println("counter: ", atomic.LoadInt64(&counter))
			wg.Done()
		}()
	}

	fmt.Println("GoRoutine: ", runtime.NumGoroutine())
	wg.Wait()

	fmt.Println("counter: ", counter)
}
