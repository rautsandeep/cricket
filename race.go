package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {

	counter := 0
	const gs = 100
	fmt.Println("counter:", counter)
	fmt.Println("gs:", gs)
	fmt.Println("Number of CPU: ", runtime.NumCPU())
	fmt.Println("GoRoutine: ", runtime.NumGoroutine())
	wg.Add(gs)
	var mu sync.Mutex
	for i := 0; i < gs; i++ {

		go func() {
			mu.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mu.Unlock()
			wg.Done()
		}()
	}

	fmt.Println("GoRoutine: ", runtime.NumGoroutine())
	wg.Wait()
	
	fmt.Println("counter: ", counter)
}
