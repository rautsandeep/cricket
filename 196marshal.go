package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string
	Age  int
}

func main() {

	u1 := user{
		Name: "sandeep",
		Age:  23,
	}
	u2 := user{
		Name: "snehal",
		Age:  22,
	}
	u3 := user{
		Name: "Kumar",
		Age:  17,
	}
	u4 := user{
		Name: "Atul",
		Age:  14,
	}

	users := []user{u1, u2, u3, u4}
	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {

		fmt.Println("error: ", err)
	}

	fmt.Println(string(bs))

}
