package main

import (
	"fmt"
)

func main() {
	x:= 48

fmt.Printf("the value of the x is %v\n",x)

if x < 40 && x > 45 {
fmt.Println("x is greater than 40 and less than 45")
} else {
fmt.Println("either x is less than 40 or greater than 45")
}
}
