package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func main() {
	t := triangle{
		height: 100,
		base:   50,
	}

	s := square{
		sideLength: 60,
	}

	printArea(t, "Triangle")
	printArea(s, "Square")
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func printArea(s shape, shapeType string) {
	res := s.getArea()

	fmt.Println(shapeType, res)
}
