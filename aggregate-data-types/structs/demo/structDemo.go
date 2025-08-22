package main

import (
	"fmt"
	"strings"
)

func main() {
	type homerun struct {
		player        string
		totalHomeruns int
	}

	homerunLeaders := []homerun{
		{player: "Barry Bonds", totalHomeruns: 762},
		{player: "Hank Aaron", totalHomeruns: 755},
		{player: "Babe Ruth", totalHomeruns: 714},
		{player: "Albert Pujols", totalHomeruns: 703},
		{player: "Alex Rodriguez", totalHomeruns: 696}}

	fmt.Println("\n-----ALL TIME HOMERUN LEADERS-----")
	fmt.Println(strings.Repeat("-", 15))
	fmt.Println(homerunLeaders[0].player, homerunLeaders[0].totalHomeruns)
	fmt.Println(homerunLeaders[1].player, homerunLeaders[1].totalHomeruns)
	fmt.Println(homerunLeaders[2].player, homerunLeaders[2].totalHomeruns)
	fmt.Println(homerunLeaders[3].player, homerunLeaders[3].totalHomeruns)
	fmt.Println(homerunLeaders[4].player, homerunLeaders[4].totalHomeruns)
}
