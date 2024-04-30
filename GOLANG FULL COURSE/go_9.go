package main

import (
	"fmt"
)

// Arrays in Golang
/*
func main() {
	// There are several types of defining arrays in Go

	// 1st case, defining with exact number of parameters in array
	// var array [5]string = {}

	// 2nd case defining array (with exact number of params)
	names := [3]string{"John", "Paul", "George"}
	fmt.Println(names)

	// 2nd case without exact number of params
	newNames := []string{"Andrew", "Mattew", "Vlad", "Peter"}
	fmt.Println(newNames, len(newNames))

	for i := 0; i < len(newNames); i++ {
		fmt.Println(newNames[i])
	}

	for el := 0; el < len(newNames); el++ {
		fmt.Println(newNames[el])
	}
}
*/

func main() {

	// Task find average:

	marks := [5]float64{5, 3.5, 4.8, 4.6, 4.2}

	totalNumbers := len(marks)

	totalSum := 0.0

	for i := 0; i < totalNumbers; i++ {
		totalSum += marks[i]
	}

	result := totalSum / float64(totalNumbers)

	fmt.Printf("The average total of all %d marks is %.1f\n", totalNumbers, result)

}
