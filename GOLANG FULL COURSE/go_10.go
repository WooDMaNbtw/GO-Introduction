package main

import "fmt"

// double arrays

func main() {
	array := [3][3][2]int32{
		{
			{213, 21},
			{92, 29},
			{84, 57}},
		{
			{57, 78},
			{82, 94},
			{23, 21}},
		{
			{23, 12},
			{78, 92},
			{23, 18}}}

	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			for k := 0; k < len(array[i][j]); k++ {
				print(array[i][j][k], " ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}
