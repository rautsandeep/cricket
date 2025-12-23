package main

import (
	"fmt"
)

type cricket struct {
	event string
	teams int
	venue string
	six   int
	four  int
}

type game struct {
	cricket
	format string
}

func (p cricket) schdule() {
	fmt.Println("the worldcup is won by India")

}
func (p game) result() {
	fmt.Println("champion tropy hosted by ", p.cricket.venue)
}

type sport interface {
	result()
}

func world(k sport) {

	fmt.Println("this is executed from inside of the world function")
	fmt.Println(k)
}

func main() {

	p1 := cricket{
		event: "worldcup",
		teams: 8,
		venue: "india",
		six:   290,
		four:  540,
	}

	p2 := game{
		cricket: cricket{
			event: "championtrophy",
			teams: 6,
			venue: "Aus",
			six:   190,
			four:  440,
		},
		format: "outdoor",
	}
	fmt.Println(p1.event)
	fmt.Println(p2.format)
	p1.schdule()
	p2.result()
	world(p2)

}
