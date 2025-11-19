// go has lexical (static) scoping - scope is calculated at the compile time

// types of scope in go - universal, package, file, block, function
// universal scope - outermost, available everywhere. ex- predeclared identifiers like int, string, const, etc

package main

import "fmt" // file scope - anywhere inside this file

var x = 10      // file scope, package scope - anything declared at the top level of a Go file belongs to the package scope and can be accessed anywhere inside package.
const Pi = 3.14 // package scope - capital so exported

func add(a int, b int) int {
	sum := a + b // funcion scope
	return sum
}

func main() {
	fmt.Println("scope and shadowing in go")

	if x > 5 {
		x := 1         // block scope (aything inside {})
		fmt.Println(x) // inner x used
		// variable shadowing - when a variable declared within a certain scope has the same name as a variable in an outer scope. The inner variable effectively “shadows” the outer variable, meaning that the inner variable is accessible in that scope while the outer variable is temporarily hidden or inaccessible.
	}
	fmt.Println(x)

	fmt.Println(add(56, 44))
	// avoiding shadowing is a good practice for writing go code

}
