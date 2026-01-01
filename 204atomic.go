package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {

	fmt.Println("CPU Count: ", runtime.NumCPU())
	fmt.Println("Number of goroutine: ", runtime.NumGoroutine())
	var counter int64
	gs := 100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < 100; i++ {

	fmt.Println("Number of goroutine: ", runtime.NumGoroutine())
		go func() {
atomic.AddInt64(&counter,1)
			runtime.Gosched()
		fmt.Println(	atomic.LoadInt64(&counter))
			wg.Done()
		}()
//		fmt.Println("counter:", counter)
	}

	wg.Wait()

	fmt.Println("Number of goroutine: ", runtime.NumGoroutine())
	fmt.Println("counter:", counter)

}
