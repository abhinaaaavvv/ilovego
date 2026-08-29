package main

import (
	"fmt"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[1:3]
	fmt.Println(arr)
	fmt.Println(slice)

	q := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(q)

	r := []bool{true, false, false, true, false, true, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, false},
		{q[2], r[2]},
	}
	fmt.Println(s)

	t := q[3:]
	fmt.Println(t)

	u := q[:3]
	fmt.Println(u)

	v := q[:]
	fmt.Println(v)

	a := make([]int, 5)
	b := make([]int, 5, 10)
	fmt.Println(a, b)

}
