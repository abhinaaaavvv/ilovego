package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Print("hello ")

	for i := range 10 {
		defer fmt.Println(i)
	}
}
