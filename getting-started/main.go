package main //<--- Package identifier

// fmt package allows us to work with strings
import "fmt"

// Need to have entry point. This is a function same name as package identifier
func main() {
	// println("Hello world!")
	//Using fmt uses a more powerful function but we have to take on another dependency
	fmt.Println("Hello world!")
}
