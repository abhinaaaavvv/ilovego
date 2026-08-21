package main

import (
	"fmt"

	"github.com/abhinaaaavvv/ilovego/pubilc/adder"
)

func add(a int, b int) (int, int) {
	return a + b, a * b
}

func main() {
	a := 5
	b := 10
	if a > b {
		if b > a {
			println(a)
		}
	}

	sum1 := adder.Adder(a, b)

	sum, product := add(a, b)
	fmt.Println("the sum of a and b is", sum)
	fmt.Println("the product of a and b is", product)
}
