package main

import (
	"fmt"
)

func main() {

	xsnehal := []string{"snehal", "shinde", "apple", "limbusarbat"}
	xsandeep := []string{"sandeep", "raut", "pear", "tea"}

	fmt.Println(xsnehal)
	fmt.Println(xsandeep)

	xxs:= [][]string{xsnehal, xsandeep}
	fmt.Println(xxs)
	fmt.Println(len(xxs))
	fmt.Println(cap(xxs))
}
