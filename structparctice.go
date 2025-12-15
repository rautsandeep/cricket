package main

import (
	"fmt"
)

type dasara struct {
	sweet    string
	mixture  string
	quantity int
}
type diwali struct {
	dasara
	poli string
}

func main() {
	d := dasara{

		sweet:    "ladu",
		mixture:  "farsan",
		quantity: 2,
	}

	diwa := diwali{
		dasara: dasara{
			sweet:    "besanladu",
			mixture:  "lasun chiwada",
			quantity: 4,
		},
		poli: "puran",
	}

	fmt.Println("The variety we have for dasara", d)
	fmt.Printf("%T\n", d)
	fmt.Printf("%#v, %v\n", d, d)

	fmt.Println(diwa.poli)
	fmt.Println(diwa.dasara.quantity, diwa.poli)
}
