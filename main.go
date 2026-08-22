package main

import (
	"fmt"
)

func main() {
	arr := [10]int{1, 0, 2, 9, 3, 8, 4, 7, 5, 6}

	sorted := sort(arr[:])

	fmt.Println(sorted)
}

func sort(arr []int) []int {
	n := len(arr)
	if n <= 1 {
		return arr
	}

	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return arr
}
