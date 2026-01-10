package main

import (
	"fmt"
)

func proAdder(values ...int) int { // variadic functions
	total := 0
	for _, val := range values {
		total += val
	}
	return total
}

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) GetAge() { // method of struct User
	fmt.Println(u.Age)
}

// function params work on copy by values by default
func (u User) SetAge() {
	u.Age = 50
}

func main() {
	fmt.Println("functions")
	// func inside func are not allowed

	fmt.Println(proAdder(1, 2, 3, 4, 5, 6, 7, 8, 9))

	var u = User{"avi", "aviral@go.dev", 24}
	u.GetAge()
	u.SetAge()
	fmt.Println(u)
}
