package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	c:=[]byte{102,33,56}
	d:= bytes.NewBuffer(c)
	fmt.Println(d)
	a:= bytes.NewBuffer([]byte("Snehal"))
	b := bytes.NewBufferString("hi my name is sandeep")
	fmt.Println(b)

	log.Println(a)
	b.WriteString("I like cricket")
	log.Println(b)

	b.String()
	fmt.Println(b)
	fmt.Printf("The length of the buffer %v\n",b.Len())
	fmt.Printf("%T\n",b)
	fmt.Printf("%T\n",c)


	fmt.Println("--------------------")
	fmt.Println(d.String())
	fmt.Println(b.String())
	fmt.Printf("the number of unused bytes in buffer %v is %v",b ,b.Available() )
}

