package main

import (
	"fmt"
)

func main() {
	// const are good for performance as they are evalulated at compile time
	fmt.Println("const and iota in go")
	const Pi = 3.14
	const (
		// can declare multiple consants, typed or untyped
		Error            = "null"
		StatusOK         = 200
		StatusError      = 500
		MaxSize     int  = 1024
		Debug       bool = true

		// Constant expressions
		Hundred = 10 * 10
		Million = Hundred * Hundred * Hundred
	)

	// iota starts at 0 in every const block, and each new line in the same const block increments iota by 1
	const (
		alpha = iota // = 0
		beta         // = 1
		gamma        // = 2
	)

	// each const block resets iota to 0
	// useful for enums - go does not have enums so iota is helpful
	const (
		_     = iota // skip 0 value
		Admin        // 1
		User         // 2
		Guest        // 3
	)
}
