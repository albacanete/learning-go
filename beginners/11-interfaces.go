package main

import (
	"fmt"
	"math"
)

func main() {
	rect := Rectangle{50, 20}
	circ := Circle{7}

	fmt.Println("Rectangle area:", getArea(rect))
	fmt.Println("Circle area:", getArea(circ))
}

// definition of Rectangle data structure
type Rectangle struct {
	height float64
	width  float64
}

// definition of Circle data structure
type Circle struct {
	radius float64
}

// common interface to rectangle and circle
type Shape interface {
	area() float64
}

// calculate area of rectangle
func (r1 Rectangle) area() float64 {
	return r1.height * r1.width
}

// calculate area of circle
func (c1 Circle) area() float64 {
	return math.Pi * math.Pow(c1.radius, 2)
}

// calculate area of any data structure
func getArea(shape Shape) float64 {
	return shape.area()
}
