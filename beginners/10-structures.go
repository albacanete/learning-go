package main

import "fmt"

func main() {
	// user defined data types
	rect1 := Rectangle{10, 5}
	// if we don't know the order of the params: Rectangle{height: 30, width: 40}
	fmt.Println("Height:", rect1.height)
	fmt.Println("Width:", rect1.width)
	fmt.Println("Area:", rect1.area())
}

type Rectangle struct {
	height float64
	width  float64
}

func (rect *Rectangle) area() float64 {
	return rect.height * rect.width
}
