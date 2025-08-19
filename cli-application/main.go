package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("What would you like me to scream?")

	// os.Stdin (standard input) allows us to get input from the keyboard 1 character at a time
	// We wrap the os.Stdin with bufio to provide some nice functionality
	input := bufio.NewReader(os.Stdin)

	// Read string allows us to read the input up to a certain point. Ex: ('\n')
	// The underscore is used to catch an error (we leave this unhandled)
	s, _ := input.ReadString('\n')
	
	// Then we trim the space and set the input to upper case
	s = strings.TrimSpace(s)
	s = strings.ToUpper(s)

	// Then we print out the input 
	fmt.Println(s)
}