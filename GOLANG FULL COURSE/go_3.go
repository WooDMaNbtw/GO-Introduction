package main

import (
	"fmt"
	"math"
)

// CONDITIONAL STATEMENTS

// if/else/else if

func main() {
	num := 3

	if num > 0 {
		fmt.Println("Number is greater than zero")
	} else if num == 0 {
		fmt.Println("Number is zero")
	} else {
		fmt.Println("Number is smaller than zero")
	}

	fmt.Println(" ----- Lets solve quadratic equation! ----- ")
	var (
		a float64
		b float64
		c float64
	)

	fmt.Println("Input a: ")
	fmt.Scanln(&a)

	fmt.Println("Input b: ")
	fmt.Scanln(&b)

	fmt.Println("Input c: ")
	fmt.Scanln(&c)

	discriminant := (b * b) - 4*(a*c)

	if discriminant > 0 {

		fmt.Println("Equation has 2 roots")
		x1 := (-b + math.Sqrt(discriminant)) / (2 * a)
		x2 := (-b - math.Sqrt(discriminant)) / (2 * a)
		fmt.Println("x1:", x1)
		fmt.Println("x2:", x2)

	} else if discriminant == 0 {
		fmt.Println("Equation has 1 root")
		x1 := -b / (2 * a)
		fmt.Println("x1:", x1)
	} else {
		fmt.Println("Equation only have complex roots")
	}

	fmt.Scanln(&a)
}
