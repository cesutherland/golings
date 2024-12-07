// interfaces1
// Make me compile!

package main

import (
	"fmt"
	"math"
)

// add a Shape interface that both Circle and Rectangle satisfy
// and is used in the CalculateArea function
type Shape interface { // something is missing here
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 { // we might need a method here to implement the Shape interface
	return c.Radius * c.Radius * math.Pi
}

func CalculateArea(shape Shape) float64 {
	return shape.Area()
}

func main() {
	circle := Circle{10}
	rectangle := Rectangle{5, 3}

	fmt.Println("Circle Area:", CalculateArea(circle))
	fmt.Println("Rectangle Area:", CalculateArea(rectangle))
}
