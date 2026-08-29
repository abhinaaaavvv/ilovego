package main

import "fmt"

type Vertex struct {
	X, Y int
}

func (v Vertex) add() int {
	return v.X + v.Y
}

func main() {
	v := Vertex{
		5, 6,
	}
	fmt.Println(v.add())
}
