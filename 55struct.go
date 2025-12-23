package main

import(
"fmt"
)

func main(){
type engine struct{
electric bool
}

type vehicle struct{

	engine
	style string
	model int
	doors int
	colour string
}

v1:= vehicle{

	engine: engine{electric: true,},
	style: "compact SUV",
	model: 2022,
	doors: 4,
	colour:  "red",
}

v2:= vehicle{

	engine: engine{
		electric: false,

},
style: "sedan",
model: 2025,
doors: 5,
colour: "Blue",
}
	fmt.Println(v1.style, v1.model, v1.colour)
	fmt.Println(v2.style, v2.model, v1.colour)



}
