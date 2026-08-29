package main

import (
	"fmt"
	"math"
)

func isAmstrong(n int) bool {
	originalNumber := n
	sum := 0

	numberOfDigits := len(fmt.Sprintf("%d", n))

	for n > 0 {
		digit := n % 10
		sum += int(math.Pow(float64(digit), float64(numberOfDigits)))
		n /= 10
	}

	return sum == originalNumber
}

func main() {
	var n int = 1000000
	fmt.Printf("Armstrong numbers between 1 and %v are:\n", n)
	for i := 1; i <= n; i++ {
		if isAmstrong(i) {
			fmt.Println(i)
		}
	}
}
