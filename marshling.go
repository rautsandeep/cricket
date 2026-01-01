package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type student struct {
	Name       string
	Class      string
	Rollnumber int
}

func main() {
	sandeep := student{
		Name:       "sandeep",
		Class:      "six",
		Rollnumber: 42,
	}

	snehal := student{
		Name:       "snehal",
		Class:      "seven",
		Rollnumber: 43,
	}
	school:= []student{sandeep, snehal}
	fmt.Println(school)
	bs, err := json.Marshal(school)
	if err != nil {
		log.Fatalf("error %s", err)
	}
	fmt.Println(string(bs))

}
