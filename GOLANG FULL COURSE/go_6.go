package main

import "fmt"

// Switch case

func main() {
	name := "Sung"

	switch name {
	case "John":
		fmt.Println("John")
		fallthrough //позволяет идти дальше, если case сработал
		// Работает только с switch case
	case "Mary":
		fmt.Println("Mary")
	case "Jack":
		fmt.Println("Jack")
	case "Tom":
		fmt.Println("Tom")
	default:
		fmt.Println("Not found")
	}
}
