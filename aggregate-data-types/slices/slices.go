package main

import (
	"fmt"
	"slices"
)

func main() {

	basicSlices()
	copyingSlices()

}

func basicSlices() {
	fmt.Println("\n---BASIC SLICES---")
	//slices of ints
	var s []int
	//This prints [], but the real value is nil
	fmt.Println("Empty slice:", s)

	// Slice literal
	s = []int{1, 2, 3}
	fmt.Println("Slice index of 1:", s[1])
	//Update the value
	s[1] = 99
	fmt.Println("Slice with updated values:", s)

	//adding elements to the slice
	s = append(s, 5, 10, 15) //<--- append will add new values to the end
	fmt.Println("Slice with added values:", s)

	// Deleting indicies 1 and 2 from the slice
	// This will delete s[1](value=99) and s[2](value=3)
	s = slices.Delete(s, 1, 3) //<--- params = (sliceName, starting index we want to rmv(1), up to but not including the last index we want to rmv(3))
	fmt.Println("Slice with deleted values:", s)
}

func copyingSlices() {
	fmt.Println("\n---COPYING SLICES---")
	// OG slice
	s := []string{"Mike,", "Dave,", "Robert"}
	// slices are copied by reference
	// use slices.Clone to clone
	s2 := s
	fmt.Println("the s2 copy of the OG slice:", s2)

	//updating values in both slices at the same time
	s[0], s[2] = "Freddy,", "Boogie Man"
	// The values of indicies 0 and 2 will change in BOTH s and s2
	fmt.Println("After changing values to only OG slice:")
	fmt.Println("\t- The OG slice:", s)
	fmt.Println("\t- The copy slice:", s2)

	// fmt.Println(s == s2) <--- This will give us a compile time error, slices are not comprable
}
