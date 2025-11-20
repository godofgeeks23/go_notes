package main

import (
	"fmt"
)

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("struct in go")

	// go does not have classes, so no inheritance/super/parent

	avi := User{"Avi", "aviralji4@gmail.com", true, 24}
	fmt.Println(avi)
	fmt.Println(avi.Name)
}
