/*

execution times: interpreted lang > go (due to go-runtime), VM based lang(java, c#) > natively compiled lang
compilation time: go > compiled lang

- strongly and statically typed
- garbage collected
- not a pure OOP language

Basic data types in go -
bool
string

int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr
byte // alias for uint8
rune // alias for int32, represents a Unicode code point
float32 float64
complex64 complex128

*/

package main // every go program is made up of packages

import (
	"fmt"
	"math"
	"math/rand"
)

// variables in go are always initialized to their zero value if not explicitly initialized
var c, python, java bool // var declares a list of variables in an scope

var k, z int = 8, 10 // this is valid, k is initialized to 8
// k:= 8 - this is not valid in package scope. can only be done in function scope

const Pi = 3.14

func add(x int, y int) int {
	return x + y
}

func swap(x int, y int) (int, int) {
	return y, x
}

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("My favorite number is", rand.Intn(10))
	i := 42
	fmt.Println(i, c, python, java)
	fmt.Println("pi =", math.Pi) // a name is exported if it begins with a capital letter

	x := 12
	y := 78
	fmt.Println("sum =", add(x, y))
	fmt.Printf("%v, %v\n", x, y)
	x, y = swap(x, y)
	fmt.Printf("%v, %v\n", x, y)

	var f float64 = float64(i) // The expression T(v) converts the value v to the type T
	fmt.Println(f)
}

// var i int - use this when initial value does not matter
// i := 42 - use this when initial value is known
