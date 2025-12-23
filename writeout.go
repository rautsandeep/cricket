package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"bytes"
)

type person struct {
	name string
}

func (p person) writing(s io.Writer) (int, error) {
	x, err := s.Write([]byte(p.name))
	return x, err
}
func main() {
	p1 := person{
		name: "snehal",
	}
	f1, err := os.Create("newfile.txt")
	if err != nil {
		log.Fatalf("error %s\n", err)
	}
	defer f1.Close()

	v, err := p1.writing(f1)
	log.Println(v)
	var z bytes.Buffer
	fmt.Printf("%T\n", z)
	count, err := p1.writing(&z)
	fmt.Println("the number of byter written is ",count)
	log.Println(v)

}
