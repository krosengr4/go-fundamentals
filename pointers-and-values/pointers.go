package main

import "fmt"

func main() {
	s := "Hello world!"

	// Using p as a pointer to s
	p := &s 
	// Printing out our pointer will print a memory address (ex: 0x14000010070)
	fmt.Println(p)
	// We can print the value of the pointer by using a dereferencing operator
	fmt.Println(*p)

	// We can use the new function
	p = new(string)
	// The memory address of p has now changed, and we now have no access to the value of p
	fmt.Println(p, *p)
}