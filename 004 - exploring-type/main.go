package main

import (
	"fmt"
)

// Go is STATICALLY typed
var y = 42
var z = "Shaken, not stirred"
var a = `James said, 
	"Shaken, 
	not stirred"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
}
