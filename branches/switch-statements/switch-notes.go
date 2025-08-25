package main

import "fmt"

func main() {

	// Using a switch statement and initializer.
	switch i := 5; i {
	case 1:
		fmt.Println("First case.")
	case 2 + 3, (2 * i) + 3: //<-- Case 2 + 3 OR case (2 * i) + 3
		fmt.Println("Second case.")
	default:
		fmt.Println("Default case.")
	}

}
