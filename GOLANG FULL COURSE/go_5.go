package main

import "fmt"

// LOOPS

/*
GoLang has only FOR loop (no while, no do while)
*/

func main() {
	/* for i := 0; i < 100; i++ {
		if i%3 == 0 || i%5 == 0 {
			fmt.Println("Aliquot by 3 or 5:", i)
		} else if i == 48 {
			break
		} else if i == 29 {
			fmt.Println("29")
			continue
		}
	}
	*/

	/*
		nums := []int{2, 3, 4, 5}
		for index, element := range nums {
			fmt.Println("Index:", index, "Element:", element)
		}
	*/

	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	for _, row := range matrix {
		for _, col := range row {
			fmt.Print(col, " ")
		}
		fmt.Println()
	}
}
