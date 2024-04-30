package main

import (
	"fmt"
)

// Struct methods

type User struct {
	name     string
	age      int
	password string
	score    []int
}

func main() {
	user := User{"John", 18, "123456", []int{29, 32, 68, 79, 83, 92, 95, 100, 78}}
	fmt.Println(user.name)
	fmt.Println(user.getName())
	user.setName("Kate")
	fmt.Println(user.getName())
	fmt.Println(user.isAdult())

	fmt.Println(user.getHighestScore())
	fmt.Println(user.sortScore())
}

func (u *User) getName() string {
	return u.name
}

func (u *User) setName(name string) {
	u.name = name
}

func (u *User) isAdult() bool {
	return u.age >= 18
}

func (u *User) getHighestScore() int {
	// sorting

	scores := u.score
	maxScore := 0

	for i := 1; i < len(scores); i++ {
		if scores[i] > maxScore {
			maxScore = scores[i]
		}
	}
	return maxScore
}

func (u *User) sortScore() []int {

	length := len(u.score)

	for i := 0; i < length-1; i++ {
		if u.score[i] > u.score[i+1] {
			u.score[i], u.score[i+1] = u.score[i+1], u.score[i]
		}
	}
	return u.score
}
