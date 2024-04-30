package main

// Files in GO

import (
	"fmt"
	"os"
)

func main() {
	create, err := os.Create("example.txt")

	if err != nil {
		fmt.Println("Couldn't create file", err)
	}

	create.WriteString("Hello World")

	fileData, err := os.ReadFile("example.txt")

	if err != nil {
		fmt.Println("Couldn't read file\n", err)
	}

	fmt.Println(string(fileData))

	os.Remove("example.txt")
}
