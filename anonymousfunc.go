package main

import(
"fmt"
)

func sum(a,b int)int{
return a+b
}
func main(){

	x:= sum(4,5)
	fmt.Println("Sum of two numbers is", x)

	d:= func(a int ,b int)(int){
return a-b
	}(10,4)
fmt.Println("difference of two number is ",d)
}
