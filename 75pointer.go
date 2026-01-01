package main

import (
	"fmt"
)

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The meaning of life is ..."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n
}

func main() {
	var p, q, r string
	var n int
	fmt.Printf("the value stored at the p is  %v and type is %T\n", p, p)
	fmt.Printf("the value stored at the q is %v and type is %T\n", q, q)
	fmt.Printf("the value stored at the r is %v and type is %T\n", r, r)
	fmt.Printf("the value stored at the n is  %v and type is %T\n", n, n)
	fmt.Printf("the value stored at the a is %v and type is %T\n", a, a)
	fmt.Printf("the value stored at the b %v and type is %T\n", b, b)
	fmt.Printf("the value stored at the c is %v and type is %T\n", c, c)
	fmt.Printf("the value stored at the d is %v and type is %T\n", d, d)

	fmt.Printf("The value stroed at the memory location %v is %v\n", &a, *a)
	fmt.Printf("The value stroed at the memory location %v is %v\n", &a, *b )
	fmt.Printf("The value stroed at the memory location %v is %v\n", &a,  *c )
	fmt.Printf("The value stroed at the memory location %v is %v\n",&d, *d)
}
