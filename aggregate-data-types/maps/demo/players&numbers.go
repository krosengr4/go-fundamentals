package main

import "fmt"

func main() {
	// Create an array list of string with our baseball players
	baseballPlayers := []string{
		"Jeter, Derek",
		"Wright, David",
		"Hamilton, Josh",
		"Rodriguez, Ivan('pudge')", //<--- Even though there is nothing after, this must end with a comma
	}

	//Create map called numbers to map the number to the player. key = string, value = int
	numbers := map[string]int{
		baseballPlayers[0]: 2,
		baseballPlayers[1]: 5,
		baseballPlayers[2]: 32,
		baseballPlayers[3]: 7, //<--- Even though there is nothing after, this must end with a comma
	}

	// Get the sum of all the players numbers
	sumOfNumbers := numbers[baseballPlayers[0]] + numbers[baseballPlayers[1]] + numbers[baseballPlayers[2]] + numbers[baseballPlayers[3]]

	fmt.Println("\n-----BASEBALL PLAYERS AND THEIR NUMBERS------")
	fmt.Println(baseballPlayers[0]+":", numbers[baseballPlayers[0]])
	fmt.Println(baseballPlayers[1]+":", numbers[baseballPlayers[1]])
	fmt.Println(baseballPlayers[2]+":", numbers[baseballPlayers[2]])
	fmt.Println(baseballPlayers[3]+":", numbers[baseballPlayers[3]])
	fmt.Println("Sum of all numbers:", sumOfNumbers)

}
