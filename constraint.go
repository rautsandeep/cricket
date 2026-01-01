package main

import(
"fmt"
// "golang.org/x/exp/constraints"
)

type num interface {
 ~int | ~float64
}
func add[T num](a,b T)T{
return a+b
}

func main(){
	type game int
	type dame float64
	var x game = 5
	var y game = 60
	var p dame = 50.34
	var q dame = 40.34
fmt.Println("addition of two numbers are ",add(x, y))
fmt.Println("addition of two numbers are ",add(p, q))

}
