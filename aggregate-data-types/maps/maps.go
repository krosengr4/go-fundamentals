package main

import "fmt"

func main() {
	basicMap()
}

func basicMap() {
	// declaring a map
	var m map[string]int //<--- key = string, value = int
	// because we just declared and didn't instantiate, the value will be nil
	fmt.Println("\nEmpty map:", m)
	// map literal
	m = map[string]int{"foo": 1, "bar": 2} //<--- key = string, value = int
	fmt.Println("Populated map:", m)

	// looking up values in a map
	fmt.Println("The value of foo is:", m["foo"])
	// updating value in the map
	m["bar"] = 800815
	fmt.Println("Updated map:", m) //<- notice that the ordering of the output changed

	// Adding a value to the map by stating key that isn't in the map yet
	m["crawdaddy"] = 194
	fmt.Println("The value of the new key, crawdaddy, is:", m["crawdaddy"])

	// Deleting a KVP from the map
	delete(m, "foo")
	fmt.Println("After deleting KVP from map:", m)
}
