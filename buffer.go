package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	a := "India"
	b := bytes.NewBufferString(a)
	fmt.Println(b)
	x, err := b.WriteString(" Is my country")
	if err != nil {
		log.Fatal("error %s", err)
	}
	fmt.Println("count of bytes written is", x)
	fmt.Println(b)
	b.Reset()
	fmt.Println(b)

	b.WriteString("India won by 8 Wickets")
	fmt.Println(b)
}
