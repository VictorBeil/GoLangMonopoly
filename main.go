package main

import (
	"fmt"
	"monopolyGO/monopoly"
)

func main() {
	var playerNames []string
	fmt.Print("Player Names: ")
	for i := 0; i < 4; i++ {
		fmt.Scan(playerNames[i])
	}
	monopoly.NewGame(playerNames)
	game := monopoly.NewGame(playerNames)

	// Hauptspielschleife (kann weiter entwickelt werden)
	for !monopoly.GameIsOver(game) {
		game.PlayTurn()
	}

	fmt.Println("Spiel beendet!")
}
