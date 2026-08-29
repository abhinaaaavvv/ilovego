package main

import (
	"fmt"
)

func main() {
	slice_1 := []int{1, 2, 3, 4, 5}
	slice_2 := []int{6, 7, 8, 9, 0}

	slice_1 = append(slice_1, slice_2...)
	fmt.Println(slice_1)
}
