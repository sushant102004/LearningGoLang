package main

import "fmt"

type Shapes interface {
	area() float64
}

type Circle struct {
	radius float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (circle Circle) area() float64 {
	return 3.14 * (circle.radius * circle.radius)
}

func (rectangle Rectangle) area() float64 {
	return 2 * (rectangle.width + rectangle.height)
}

func main() {
	circle := Circle{9.9}
	rectangle := Rectangle{3, 4}

	shapes := []Shapes{circle, rectangle}

	for _, shape := range shapes {
		fmt.Println(shape.area())
	}
}
