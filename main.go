package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var sum int = 0
	for i := range 10 {
		sum += i
	}

	if sum < 100 && sum > 0 { // expression
		fmt.Println(1)
	} else if sum == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(-1)
	}

	if product := sum * 2; product > sum { //expression; condition
		fmt.Println(product)
	}

	fmt.Print("Go is running in: ")
	switch os := runtime.GOOS; os { //expression; condition - same as if
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}

	var t = time.Now()

	switch {
	case t.Hour() < 12: //expression
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
