package main

import "fmt"

func main() {
	basicArray()
	referenceArray()

}

func basicArray() {
	fmt.Println("\n-----BASIC ARRAY-----")
	var arr [3]string //<--- array of 3 strings
	fmt.Println("Empty array:", arr)
	arr = [3]string{"Mike", "Jon", "Pete"} //<--- Array literal
	fmt.Println("Populated array:", arr)

	arr[1] = "Kevin"
	fmt.Println("Updated array:", arr)
	fmt.Println("Length of the array:", len(arr))
}

func referenceArray() {
	fmt.Println("\n-----USING ARRAYS AS A POINTER-----")
	arr := [3]string{"foo", "bar", "fuzz"}
	arr2 := arr
	fmt.Println("arr2 referencing arr", arr2)

	arr[0] = "quux"
	//Changing 1st value in arr WILL NOT change the value of arr2
	// arr and arr2 are INDEPENDENT VARIABLES. They have their own memory and store their data seperately.
	fmt.Println("Updated arr:", arr)
	fmt.Println("arr2 after updating arr:", arr2)

	// These two arrays are comprable, they can be compared
	fmt.Println("\nAre the two arrays equal?", arr == arr2)
}
