package main

import "fmt"

func main() {
	basicStruct()
	customType()
}

func basicStruct() {
	var a struct {
		name string
		age  int
	}

	fmt.Println("Empty struct:\n", a) //{ 0}

	a.name = "Mike"
	a.age = 102
	fmt.Println("Struct after giving values:\n", a) //{Mike 102}
}

func customType() {
	// Creating a custom type based on struct. keyword = type
	type myStruct struct {
		name string
		age  int
	}

	// Assign value with a struct literal
	var a = myStruct{
		name: "Jasper",
		age:  508}

	fmt.Println("Original struct:\n", a)

	// Copy the value of a into a2 (structs = value type)
	a2 := a
	// Make a change in new struct (a2)
	a2.name = "Jaff" //<--- The value of a will not change

	fmt.Println("Original struct(a) after change to a2:\n", a)
	fmt.Println("a2 after change:\n", a2)

}
