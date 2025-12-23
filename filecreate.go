package main

import (
	"fmt"
	"os"
	"log"
)

func main() {

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatalf("error %s", err)
	}

	defer f.Close()

	a := []byte("hello, this is Sandeep")

	x, err := f.Write(a)
	if err != nil {
		log.Fatalf("error %s", err)
	}
	fmt.Println("then number of bytes written in file is ",x)

}
