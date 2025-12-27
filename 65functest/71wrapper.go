package main

import (
	"fmt"
	"log"
)

func main() {

	wrapper(inner)
}

func wrapper(f func()) {
	log.Println("start")
	f()

	log.Println("end")
}

func inner() {
	fmt.Println("this is wrapped funciton")
}
