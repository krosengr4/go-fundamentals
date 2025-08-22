package main

import "fmt"

func main() {
	basicStruct()
}

func basicStruct() {
	var a struct {
		name string
		age  int
	}

	fmt.Println(a) //{ 0}

	a.name = "Mike"
	a.age = 102
	fmt.Println(a) //{Mike 102}
}
