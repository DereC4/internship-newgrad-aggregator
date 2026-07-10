package main

import "fmt"
// fmt is the standard library package for formatted I/O.

func main()  {
	// no semicolons in golang
	fmt.Println("Hello World!")

	// keyword first, type comes last
	var x int = 3
	var defaultInt int       // 0
    var defaultString string // ""
    var defaultBool bool     // false

	// shorthand variable declaration
	urmom := "Go"

	// Unused Variables are Compiler Errors
	// No Implicit Type Conversion, must cast to use different types
}