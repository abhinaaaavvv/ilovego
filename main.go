package main

import (
	"fmt"
)

func main() {
	var sum int = 0
	for i := range 10 {
		sum += i
	}
	fmt.Println(sum)
}
