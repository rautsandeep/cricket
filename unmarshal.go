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
	school := []student{sandeep, snehal}
	fmt.Println(school)
	bs, err := json.Marshal(school)
	if err != nil {
		log.Fatalf("error %s", err)
	}
	fmt.Println(string(bs))
	j := `[{"Name":"sandeep","Class":"six","Rollnumber":42},{"Name":"snehal","Class":"seven","Rollnumber":43}]`
	jb := []byte(j)
	var class []student
	err = json.Unmarshal(jb, &class)
	if err != nil {
		log.Fatalf("error %s\n", err)
	}
	fmt.Println(class)

	for _, v := range class {
		fmt.Println(v)
	}
}
