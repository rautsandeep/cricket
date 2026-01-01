package main

import (
	"fmt"
)

type outdoor struct {
	name        string
	playercount int
}

func (o outdoor) gamename() {
	fmt.Println("name of the outdoor game is ", o.name)
}

func (i *indoor) gamename() {
	fmt.Println("name of the outdoor game is ", i.name)
}

func (o outdoor) count() {
	fmt.Printf("count of the player for the outdoor game %v is %v\n", o.name, o.playercount)
}

func (i *indoor) count() {
	fmt.Printf("count of the player for the indoor game %v is %v\n", i.name, i.playercount)
}
type indoor struct {
	name        string
	playercount int
}

func (i *indoor) defi() {
	fmt.Printf("name of the indoor function is %v and player count is %v\n", i.name, i.playercount)
}
func (o *outdoor) defi() {
	fmt.Printf("name of the outdoor function is %v and player count is %v\n", o.name, o.playercount)
}

type game interface {
	gamename()
	count()
	defi()
}

func gamefunc(g game) {

	g.gamename()
	g.count()
	g.defi()
}
func main() {
	o1 := outdoor{
		name:        "cricket",
		playercount: 11,
	}

	o2 := outdoor{
		name:        "kabbadi",
		playercount: 8,
	}

	i1 := indoor{
		name:        "chess",
		playercount: 1,
	}
/*
	o1.gamename()
	o2.count()
	i1.defi()
*/
	fmt.Println("============================")
	gamefunc(&o1)
	gamefunc(&o2)
	gamefunc(&i1)

}
