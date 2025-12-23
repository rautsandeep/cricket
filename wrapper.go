package main

import (
	"fmt"
	"log"
	"strconv"
)

type city struct {
	name       string
	road       string
	temprature int
}

func (c city) String() string {

	return fmt.Sprintf("name of the city is %s situated on the road %s having temp %s", c.name, c.road, strconv.Itoa(int(c.temprature)))
}

type district struct {
	city
	name string
}

func (d district) String() string {

	return fmt.Sprintf("name of the city is %s situated on the road %s having temp %s belongs to the district %s", d.city.name, d.city.road, strconv.Itoa(int(d.city.temprature)), d.name)
}

func loginfo(s fmt.Stringer){
log.Println("Log prefixed by the wrapper function", s.String())
}

func main() {
	c1 := city{

		name:       "naviMumbai",
		road:       "thakur marg",
		temprature: 28,
	}

	d1 := district{
		city: city{

			name:       "mumbai",
			road:       "Gandhi marg",
			temprature: 29,
		},
		name: "thane",
	}

	fmt.Println(c1)
	fmt.Println(d1)
	log.Println(c1)
	log.Println(d1)

	loginfo(c1)
	loginfo(d1)
}
