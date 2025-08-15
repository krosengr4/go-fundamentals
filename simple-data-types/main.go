package main

import "fmt"

// entry point
func main() {
	// ALL LOCAL VARIABLES MUST HAVE A USAGE!
	// Declaring a variable and not using it is a compile time error in Go.
	// This makes sure that our code stays as clean as possible

	var a string = "foo"
	fmt.Println(a)

	//inferred data type
	var b = 99
	fmt.Println(b)

	// d could be either a float32 or float64 data type
	// For the short declaration syntax below, Go will use a default data type
	// For a decimal, default = float64
	// For a whole number, default = int
	d := 3.14142
	fmt.Println(d)

	var e int = int(d)
	fmt.Println(e)
	fmt.Println(e + e)
}
