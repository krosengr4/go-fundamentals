package main

import (
	"fmt"
	"strings"
)

func main() {
	//Create a splice of baseball players
	players := []string{
		"Dent, Bucky",
		"Jeter, Derek",
		"Schmidt, Mike",
	}

	jerseyNumbers := []int{20, 2, 20}

	fmt.Println("\n---PLAYERS AND THEIR NUMBERS---")
	fmt.Println(strings.Repeat("-", 14))
	fmt.Println(players[0], jerseyNumbers[0])
	fmt.Println(players[1], jerseyNumbers[1])
	fmt.Println(players[2], jerseyNumbers[2])
}
