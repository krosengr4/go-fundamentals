package main

import "fmt"

func main() {
	fmt.Println("\t---CONSTANTS---")
	constants()
	fmt.Println("-----------------------------")

	fmt.Println("\t---CONSTANT GROUPS---")
	constantGroups()
	fmt.Println("-----------------------------")

	fmt.Println("\t---IOTA---")
	iotaFunction()
	fmt.Println("-----------------------------")
}

func constants() {
	const a = 42
	const b float32 = 3

	var i int = a
	var f float64 = float64(b) //Must use type conversion if the const is explicitly stated

	fmt.Println(i, f)
}

func constantGroups() {
	const (
		d = 2 * 5
		e
	)

	fmt.Println("Constant d =", d)
	fmt.Println("Constant e =", e)
}

func iotaFunction() {
	// iota is related to the position in the constant group (0 indexed)
	const (
		a = 27 + 2
		b = iota     //1
		c = iota * 2 //2 * 2 = 4

	)

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("c =", c)
}
