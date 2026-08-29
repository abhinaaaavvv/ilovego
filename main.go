package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func scaleFunc(v *Vertex, f float64) {

	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{5, 6}
	v.scale(5)
	scaleFunc(&v, 10)

	p := &Vertex{7, 8}
	p.scale(5)
	scaleFunc(p, 10)

	fmt.Println(v, p)

}
