package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	perimeter() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

type Square struct {
	height, width float64
}

func (s Square) area() float64 {
	return s.height * s.width
}

func (s Square) perimeter() float64 {
	return (2 * s.height) + (2 * s.width)
}

func Area(s Shape) float64 {
	return s.area()
}

func Perimeter(s Shape) float64 {
	return s.perimeter()
}

func main() {
	a := Circle{5}
	fmt.Println(Area(a))
	fmt.Println(Perimeter(a))

	b := Square{5, 10}
	fmt.Println(Area(b))
	fmt.Println(Perimeter(b))
}
