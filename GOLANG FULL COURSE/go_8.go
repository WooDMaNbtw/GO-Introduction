package main

import "fmt"

// Mathematics operations

func main() {
	var (
		operator string
		num1     int16
		num2     int16
		result   int16
	)

	println("--- Welcome to Terminal Calculator ---")
	println("Please enter the operator: ")
	fmt.Scan(&operator)

	println("Please enter the num1: ")
	fmt.Scan(&num1)
	println("Please enter the num2: ")
	fmt.Scan(&num2)

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	case "%":
		result = num1 % num2
	}
	println("Computed result is:", result)
}
