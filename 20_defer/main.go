package main

import (
	"fmt"
)

// defer statement in a functions are executed at the end of the function
// defer statements inside a function are executed in rev order

func main() {
	defer fmt.Println("defer 1") // will be executed at the end of this function
	fmt.Println("normal line 1")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	fmt.Println("normal line 2")
	fmt.Println("normal line 3")

}
