package main

import "fmt"

type numbers interface {
	int | float64 | uint
}

func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		if v == x {
			return i
		}
	}
	return -1
}

func Add[N numbers](a ...N) N { //variadic function (a ...N)
	var sum N = 0

	for _, v := range a {
		sum += v
	}
	return sum
}

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(Index(s1, 5))

	s2 := []string{"abhinav", "moni", "bhavani"}
	fmt.Println(Index(s2, "hello"))

	fmt.Println(Add(5, 6))
	fmt.Println(Add(5.5, 6.7, 9.2))
	fmt.Println(Add(-5, 6, 14))
}
