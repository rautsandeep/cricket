package main

import (
	"encoding/json"
	"fmt"
)

type poem struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func main() {

	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`

	fmt.Println(s)
	var unmarshaled []poem
	err := json.Unmarshal([]byte(s), &unmarshaled)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(unmarshaled)

	for _,v:= range unmarshaled{
		fmt.Println(v)
		fmt.Println("---------------")
		fmt.Println("name:",v.First,"\tLastname:", v.Last,"\tAge:", v.Age,"\tSayings", v.Sayings)
		for indes,value:= range v.Sayings{
			fmt.Println("index: ",indes, "\tvalue:",value)
		}
	}
}
