package main

import (
	"fmt"
)

func main() {
	fmt.Println("pointers in go")

	a := 10
	ptr := &a
	fmt.Println(a, &a, ptr, *ptr)

	var str string = "abc"
	var strptr *string = &str
	fmt.Println("str:", str, ", &str:", &str, ", strptr:", strptr, ", *strptr:", *strptr)
}
