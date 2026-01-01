package main

import (
	"fmt"
	"sort"

)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}
type byage []user

func (a byage) Len() (int) {
        return len(a)
}

// Swap is part of sort.Interface.
func (a byage) Swap(i, j int) {
        a[i], a[j] = a[j], a[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (a byage) Less(i, j int) bool {
        return a[i].Age < a[j].Age
}

type byfirst []user

func (x byfirst) Len() (int) {
        return len(x)
}

// Swap is part of sort.Interface.
func (x byfirst) Swap(i, j int) {
        x[i], x[j] = x[j], x[i]
}

// Less is part of sort.Interface. It is implemented by calling the "by" closure in the sorter.
func (x byfirst) Less(i, j int) bool {
        return x[i].First < x[j].First
}



func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)


sort.Sort(byage(users))
fmt.Println(users)

for i,v:= range users{
	fmt.Println("index number: ",i)
	fmt.Println("\t \t Name:",v.First," ",v.Last)
	fmt.Println("\t \t Age: ",v.Age)
	sort.Strings(v.Sayings)
	for index,value:= range v.Sayings{
		fmt.Printf("\t   %v %v\n",index, value)
	}
}

fmt.Println("--------------------------")
fmt.Println("--------------------------")
fmt.Println("--------------------------")
sort.Sort(byfirst(users))

for i,v:= range users{
        fmt.Println("index number: ",i)
        fmt.Println("\t \t Name:",v.First," ",v.Last)
        fmt.Println("\t \t Age: ",v.Age)
        for index,value:= range v.Sayings{
                fmt.Printf("\t   %v %v\n",index, value)
        }
}


}

