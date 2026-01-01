package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var x sync.WaitGroup

func foo() {
	for i := 0; i < 10; i++ {

		fmt.Println("foo: ", i)
	}
	time.Sleep(15 * time.Second)
	x.Done()

}

func bar() {
	for j := 0; j < 10; j++ {
		fmt.Println("bar: ", j)
	}
}
func main() {

	fmt.Println("We are checking the parallelism")

	fmt.Println("GOOS\t\t", runtime.GOOS)
	fmt.Println("GOARCH\t\t", runtime.GOARCH)
	fmt.Println("NumCPU\t\t", runtime.NumCPU())
	fmt.Println("GORoutine\t\t", runtime.NumGoroutine())

	x.Add(1)
	go foo()
	bar()
	fmt.Println("---------------------------\n")
	fmt.Println("GOOS\t\t", runtime.GOOS)
	fmt.Println("GOARCH\t\t", runtime.GOOS)
	fmt.Println("NumCPU\t\t", runtime.NumCPU())
	fmt.Println("GORoutine\t\t", runtime.NumGoroutine())
	x.Wait()
	fmt.Println("----------")
}
