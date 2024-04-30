package main

import "fmt"

// Comparison operators

// Logical operators

func main() {

	a := 1

	if a < 0 {
		println("a is smaller than zero")
	} else if a >= 1 {
		println("a is greater or equal than one")
	}

	x := 3
	y := 10
	var custom_var int8

	fmt.Scanln(&custom_var)

	if x > 1 && y < 5 {
		println("It works")
	} else if custom_var == 8 {
		println("Thats true")
	} else {
		println("Nothing was matched!")
	}

}
