package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func foo() {

	for i := 0; i < 10; i++ {
		fmt.Println("foo iteration count", i)
	}
	wg.Done()
}

func bar() {
	for j := 0; j < 10; j++ {
		fmt.Println("I am from bar", j)
	}
}
func main() {

	fmt.Println("main function started")
	fmt.Println("Count of CPU: ", runtime.NumCPU())
	fmt.Println("Number of routines: ", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()

	wg.Wait()
	fmt.Println("main func finished")

}
