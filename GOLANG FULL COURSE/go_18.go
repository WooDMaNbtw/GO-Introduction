package main

import "fmt"

// pointers

/*
func main()  {
	a := 5

	pointer := &a

	fmt.Println(pointer)  //ссылка
	fmt.Println(*pointer) //значение

}
*/

// 1st case changing variable using functions
/*
func main() {
	str := "LLL"
	fmt.Println(str)
	change_var(&str)
	fmt.Println(str)
}

func change_var(s *string) {
	*s = "LoL"
}
*/

func main() {
	str := "NO WAY"
	var new_string *string = &str
	fmt.Println(new_string)
	fmt.Println(str)
	change(new_string)
	fmt.Println(str)
}

func change(str *string) {
	*str = "THIS IS THE WAY"
}
