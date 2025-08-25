package main

import (
	"fmt"
	"strings"
)

func main() {
	type score struct {
		name  string
		score int
	}

	scores := []score{
		{name: "Shelby Kirby", score: 87},
		{name: "Mathieu Toussignant", score: 73},
		{name: "Alexander Litto", score: 96},
		{name: "Ashley Docker", score: 78}}

	fmt.Println("Select a score to print (1 - 4):")

	// Using the Scanln function from the fmt package we can scan the user input and save it to a variable
	var userChoice string
	fmt.Scanln(&userChoice) //<--- "Address of" operator (&) to pass the address of the userChoice variable

	// Create "index" variable to tag the index
	var index int
	// If statement to tag index variable based on the user input
	if userChoice == "1" { //<--- Should use strconv package to convert to and from strings
		index = 0
	} else if userChoice == "2" {
		index = 1
	} else if userChoice == "3" {
		index = 2
	} else if userChoice == "4" {
		index = 3
	} else {
		fmt.Println("Unknown option! Defaulting to 1")
		index = 0
	}
	fmt.Println("Student Score:")
	fmt.Println(strings.Repeat("-", 15))
	fmt.Println(scores[index].name, scores[index].score)
	fmt.Println(strings.Repeat("-", 15))
}
