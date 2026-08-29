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
	fmt.Println("Armstrong numbers between 1 and 100000000 are:")
	for i := 1; i <= 100000000; i++ {
		if isAmstrong(i) {
			fmt.Println(i)
		}
	}
}
