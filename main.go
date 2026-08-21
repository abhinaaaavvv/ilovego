package main

import (
	"fmt"
)

func main() {
	var sum int = 0
	var i = 0
	for i < 10 {
		sum += i
		i++
	}
	fmt.Println(sum)
}
