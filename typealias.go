package main

import(
"fmt"
)

type num interface{
~int | ~float64
}

func add[T num](a, b T)T{
return a+b
}

func main(){

	fmt.Println("The sum of the 7 and 123.44 is", add(7, 123.44))
	type ank int
var p ank = 42
var q ank = 5
	fmt.Println("The sum of two ank is", add(p, q))

}


