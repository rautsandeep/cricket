package main

import(

"fmt"
"sort"
)



func main(){

	xi:= []int{30,49,39,20,47,57,87454,763}
	xs:= []string{"snehal","sandeep","bhavesh","teju","shree"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)
	fmt.Println("------------------")

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

}
