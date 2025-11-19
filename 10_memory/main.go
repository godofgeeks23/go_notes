package main

import (
	"fmt"
)

// memory allocations and deallocation happens automatically in golang - automatic garabage collection

func main() {
	fmt.Println("memory management in golang")

	// new() - allocates memory but no init. returns a zeroed storage - means address can;t be used directly

	// make() - allocated memory and inits. returns a memory address which is non-zeroed so can be used directly

}
