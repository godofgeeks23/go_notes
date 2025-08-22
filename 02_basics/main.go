/*
go is fast in execution than interpreted langs
go is fast in compilation than compiled langs
go is not fast than natively compiled langs in execution
go is natively compiled but its exec speed is similar to VM Langs (java, c#) due to go runtime
go is strongly and statically typed
go is garbage collected
go is not a pure OOP language

Numeric data types in go -
- signed
- unsigned
- float
 - complex
*/

package main

import "fmt"

// variables in go are always initialized to their zero value if not explicitly initialized
var c, python, java bool

var k int = 8 // this is valid, k is initialized to 8
// k:= 8 - this is not valid in package scope. can only be done in function scope

func main() {
	fmt.Println("Hello, World!")
	i := 42
	fmt.Println(i, c, python, java)
}

// summmary:
// var i int - use this when initial value does not matter
// i := 42 - use this when initial value is known
