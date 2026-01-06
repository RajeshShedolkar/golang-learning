package main

import (
	"fmt"
	"math"
)

/*
type Shape interface
{ area() float64 }
Like a struct an interface is created using the type keyword,
followed by a name and the keyword interface.
But instead of defining fields, we define a “method set”.
A method set is a list of methods that a type
must have in order to “implement” the interface.
*/

type Shape interface {
	area() float64
}

type circle struct {
	name string
	r    float64
}

type rectangle struct {
	name string
	l, w float64
}

func main() {
	c := new(circle)
	c = &circle{"Circle", 2}
	rect := new(rectangle)
	rect = &rectangle{"Rectangle", 4, 5}
	fmt.Println(*c, *rect)
	fmt.Println("Total Area:", totalArea(c, rect))

}

func totalArea(shape ...Shape) float64 {
	total := 0.0
	for _, currShape := range shape {
		total += currShape.area()
	}
	return total
}

// define receiver
func (c *circle) area() float64 {
	area := math.Pi * c.r * c.r
	fmt.Println("CurrentShape:", c.name, area)
	return area
}

func (rect *rectangle) area() float64 {
	area := rect.l * rect.w
	fmt.Println("CurrentShape:", rect.name, area)
	return area
}
