package main

import (
	"fmt"
	"sort"
)

// slices

func main() {
	slice := []int{3, 1, 2, 4, 8, 6, 9}
	slice = append(slice, 4)
	sort.Ints(slice)
	fmt.Println(slice)

	slice = append(slice, 8)

	newString := []string{"b", "d", "w", "z", "m", "l"}

	sort.Strings(newString)
	fmt.Println(newString)

}
