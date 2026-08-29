package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	var a int = 15

	var p = &a

	var point = Vertex{
		X: 12,
		Y: 34,
	}

	var pn = &Vertex{2, 4}

	fmt.Println(point.Y)
	fmt.Println(p)
	fmt.Println(pn)
}
