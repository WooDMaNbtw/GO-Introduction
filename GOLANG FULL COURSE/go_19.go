package main

import "fmt"

// structures

type old_User struct {
	name     string
	age      int
	password string
}

func main() {

	// 1st case

	var user User = User{name: "zhangsan", age: 18, password: "123456"}
	fmt.Println(user)

	// 2nd case

	newUser := User{name: "John", age: 19, password: "123456"}
	fmt.Println(newUser)

	/// getting their properties

	fmt.Println(newUser.name)
	fmt.Println(user.name)
	fmt.Println(newUser.password)
	newUser.password = "654321"
	fmt.Println(newUser.password)

	// renaming user
	fmt.Println("--------Renaming User--------")
	fmt.Println(newUser.name)
	changeName(&newUser)
	fmt.Println(newUser.name)
}

func changeName(u *User) {
	u.name = "Kate"
}
