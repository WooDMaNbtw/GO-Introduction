package main

import "fmt"

// maps

func main() {
	money := map[string]int{
		"dollars": 1000,
		"euros":   3000,
		"apples":  3,
	}

	fmt.Println(money)

	money["dollars"] = 2000
	delete(money, "apples")
	fmt.Println(money)
	fmt.Println(money["dollars"])

	fmt.Println(money["NOWAT"]) // returns 0 as default (wasn't found)

	el, status := money["dollars"]

	fmt.Println(el, status)
}
