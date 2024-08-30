package main

import (
	"fmt"
	"math"
)

type Shape interface {
	getArea() float64
}

type Square struct{
	sideLenght float64
}

type Triangle struct{
	base float64
	height float64
}

func main() {
	square := Square{
		sideLenght: 2,
	}
	triangle := Triangle{
		base: 2,
		height: 3,
	}
	printArea(square)
	printArea(triangle)
}

func printArea(s Shape) {
	fmt.Printf("Area: %v\n", s.getArea())
}

func (s Square) getArea() float64 {
	return math.Pow(s.sideLenght, 2)
}

func (t Triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}