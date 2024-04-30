package main

import "fmt"

// functions

func main() {
	minor()
	sum(1, 3)

	sum := returnedSun(1, 2)
	fmt.Println(sum)

	sum, sub, mul, div := mathFunc(3, 6)

	fmt.Println(sum, sub, mul, div)
}

func minor() {
	fmt.Println("minor function")
}

func sum(a int, b int) {
	fmt.Println(a + b)
}

func returnedSun(a int, b int) int {
	return a + b
}

func mathFunc(a int, b int) (int, int, int, int) {
	sum := a + b
	sub := a - b
	mul := a * b
	div := a / b

	return sum, sub, mul, div
}
