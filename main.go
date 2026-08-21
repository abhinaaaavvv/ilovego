package main

import "fmt"

func add(a *int, b *int) int {
	return *a + *b
}
func main() {
	var a = 5
	var b = 10

	var sum = add(&a, &b)

	fmt.Println(sum)

}
