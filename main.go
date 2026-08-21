package main

import (
	"fmt"
	"math"

	"ilovego/public"
)

func add(a int, b int) (int, int) {
	return a + b, a * b
}

var (
	number = 1234
	name   = "abhinav"
)

const (
	Pi  = math.Pi
	age = 16
)

func main() {
	var a = 5
	var b = 10

	var sum = public.Add(a, b)
	var product = public.Multi(a, b)

	fmt.Println("the sum of a and b is", sum)
	fmt.Println("the product of a and b is", product)
}
