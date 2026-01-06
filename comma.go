package main

import (
	"fmt"
)

func main() {

	c := make(chan int, 2)

	c <- 44
	close(c)
	v, ok := <-c
	if ok {

		fmt.Println("value pulled successfully and value is ", v, ok)
	} else {
		fmt.Println("no success while pulling value", ok)
	}
	v, ok = <-c
	if ok {

		fmt.Println("value pulled successfully and value is ", v, ok)
	} else {
		fmt.Println("no success while pulling value", ok)
	}

}
