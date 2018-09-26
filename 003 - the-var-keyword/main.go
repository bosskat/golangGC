package main

import (
	"fmt"
)

// DECLARE the variable "y"
// ASSIGN the value 43
// declare & assign = initialisation
var y = 43

// DECLARE there is a VARIABLE with the IDENTIFIER "z"
// and that "z" is of TYPE int
// ASSIGN the ZERO VALUE of TYPE int to "z"
// false for booleans, 0 for integers, 0.0 for floats,
// and nil for pointers, functions, interfaces, slices, channels, and maps.
var z int

func main() {
	// short declaration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
	x := 42
	fmt.Println(x)

	fmt.Println(y)

	foo()

	fmt.Println(z)

}
func foo() {
	fmt.Println("again:", y)
}
