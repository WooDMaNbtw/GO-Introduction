package main

import "fmt"

func main() {
	a := func(x int, y int) int {
		return x + y
	}

	square := func(x int) int {
		return x * x
	}

	sum := a(2, 20)
	square_result := square(sum)

	fmt.Println(sum)
	fmt.Println(square_result)

	foo(square)
}

//создаем функцию вычисления квадрата

//Кидаем ее ссылку в foo, затем передаем в нее параметры, в итоге запись такая: square(2)

func foo(someFunc func(x int) int) {
	fmt.Println(someFunc(2))
}
