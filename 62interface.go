package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}
type Circle struct {
	radius float64
}

func (s square) areaOfSquare() float64 {
	area := s.length * s.width
	return area
}

func (c Circle) areaOfCircle() float64 {
	area := math.Pi * c.radius * c.radius
	return area
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return s.length* s.width

}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius,2)
}

func info(s shape) {
	fmt.Println("The area of the shape is ", s.area())
}
func main() {
	s1 := square{

		length: 5,
		width:  4,
	}
	s2 := Circle{
		radius: 6,
	}

	x := s1.areaOfSquare()
	fmt.Println("The are of the square is", x)

	y := s2.areaOfCircle()
	fmt.Println("The area of the circle is", y)
	info(s1)
	info(s2)
}
