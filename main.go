package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v Vertex) add() float64 {
	return v.X + v.Y
}

func (v *Vertex) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{
		5, 6,
	}
	fmt.Println(v.add())

	v.scale(5)
	fmt.Println(v.add())
}
