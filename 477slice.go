package main

import (
	"fmt"
)

func main() {
	y := make([]string, 0)
	y = append(y, " Alabama", " Alaska", " Arizona", " Arkansas", " California", " Colorado", " Connecticut", "Delaware", " Florida", " Georgia", " Hawaii", " Idaho", " Illinois", " Indiana", " Iowa", " Kansas", "Kentucky", " Louisiana", " Maine", " Maryland", " Massachusetts", " Michigan", " Minnesota", "Mississippi", " Missouri", " Montana", " Nebraska", " Nevada", " New Hampshire", " New Jersey", " New Mexico", " New York", " North Carolina", " North Dakota", " Ohio", " Oklahoma", " Oregon", " Pennsylvania", " Rhode Island", " South Carolina", " South Dakota", "Tennessee", " Texas", "Utah", " Vermont", " Virginia", " Washington", " West Virginia", " Wisconsin", " Wyoming")

	fmt.Println(y)
	fmt.Printf("the length of the slice is %v\n", len(y))
	fmt.Printf("The capacity of the slice is %v\n", cap(y))

	for i := 0; i < len(y); i++ {

		fmt.Printf("index %v  %v\n", i, y[i])
	}
}
