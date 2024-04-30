package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 9, 10}

	for index, value := range slice {
		fmt.Printf("Index: %d; Value: %d\n", index, value)
	}

	/*
		output:
		Index: 0; Value: 1
		Index: 1; Value: 2
		Index: 2; Value: 3
		Index: 3; Value: 4
		Index: 4; Value: 5
		Index: 5; Value: 6
		Index: 6; Value: 9
		Index: 7; Value: 10
	*/
}
