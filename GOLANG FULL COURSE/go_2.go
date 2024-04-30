package main

import "fmt"

func main() {

	// 1st case defining vars
	fmt.Println(" ----- Here is the 1st case defining variables ----- ")

	var age int8 = 23
	fmt.Println("Here is the integer:", age)

	var number float64 = 275.9213
	fmt.Println("Here is the float:", number)

	var name string = "Bob"
	fmt.Println("Here is the string:", name)

	// 2nd case defining vars
	fmt.Println(" ----- Here is the 2nd case defining variables ----- ")
	var (
		lastname  string
		digit     int8
		newNumber float64
	)

	fmt.Println("Here is the fullname:", lastname)
	fmt.Println("Here is the digit:", digit)
	fmt.Println("Here is the new number:", newNumber)

	// 3rd case defining vars
	fmt.Println(" ----- Here is the 3rd case defining variables ----- ")

	newName := "Bob"
	newAge := 20
	newFloatDigit := 92.2313

	fmt.Println("Here is the fullname:", newName)
	fmt.Println("Here is the digit:", newAge)
	fmt.Println("Here is the new float:", newFloatDigit)

	// Users input section:
	fmt.Println(" ----- Here is the user's input section ----- ")

	var (
		usersName     string
		usersAge      int8
		usersBirthday string
	)
	fmt.Println("Please enter your full name: ")
	fmt.Scanln(&usersName)
	fmt.Println("Please enter your age: ")
	fmt.Scanln(&usersAge)
	fmt.Println("Please enter your birth date (dd/mm/yyyy): ")
	fmt.Scanln(&usersBirthday)

	fmt.Println("We have filled the application form Mr", usersName)
	fmt.Println("Name:", usersName)
	fmt.Println("Age: " + fmt.Sprint(usersAge) + " y.o.")
	fmt.Println("Birth:", usersBirthday)

}
