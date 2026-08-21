package main

import (
	"fmt"

	"github.com/abhinaaaavvv/ilovego/public"
)

func add(a int, b int) (int, int) {
	return a + b, a * b
}

func main() {
	a := 5
	b := 10

	sum1 := public.Add(a, b)
	product1 := public.Multi(a, b)

	sum, product := add(a, b)

	fmt.Println("the sum of a and b is", sum)
	fmt.Println("the sum of a and b is", sum1)
	fmt.Println("the product of a and b is", product)
	fmt.Println("the product of a and b is", product1)
}
