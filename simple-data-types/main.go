package main

import "fmt"

// entry point
func main() {
	fmt.Println("\t---VARIABLE DECLARATION---")
	varDeclaration()
	fmt.Println("------------------------------------------")

	fmt.Println("\t---ARITHMETIC---")
	arithmetic()
	fmt.Println("------------------------------------------")

	fmt.Println("\t---COMPARISONS---")
	comparisons()
	fmt.Println("------------------------------------------")

}

func varDeclaration() {
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

func arithmetic() {
	a, b := 5, 2

	fmt.Println("First Number =", a)
	fmt.Println("Second Number =", b)
	fmt.Println("----------------------")

	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)

	// Division will return a whole number
	fmt.Println("Division:", a/b)
	fmt.Println("Remainder of division:", a%b)

	// Converting a and b to floats so result will be a float
	fmt.Println("Float division:", float32(a)/float32(b))
}

func comparisons() {
	a, b := 3, 7

	fmt.Println("First Number =", a)
	fmt.Println("Second Number =", b)
	fmt.Println("----------------------")

	fmt.Println("Are they equal?:", a == b)
	fmt.Println("Is number 1 bigger?:", a > b)
	fmt.Println("Is number 2 bigger?:", a < b)
}
