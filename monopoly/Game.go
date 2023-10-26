package monopoly

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	Players     []*Player
	Board       *Board
	CurrentTurn int
}

func (game *Game) PlayTurn() {
	var chooseInt int
	currentPlayer := game.Players[game.CurrentTurn]
	fmt.Print("---------------------Your Move---------------------------")
	fmt.Print("1 -[Throw dices]")
	fmt.Print("2 -[Sell Property]")
	fmt.Print("3 -[ViewPlayer]")
	fmt.Scan(chooseInt)
	switch chooseInt {
	case 1:

	case 2:

	case 3:
		fmt.Print(currentPlayer.properties)
	default:
		fmt.Println("No Player_Input")
	}
	// Aktualisiere den Spielzustand (z.B. aktuellen Spieler wechseln)
	game.CurrentTurn = (game.CurrentTurn + 1) % len(game.Players)
}

func NewGame(playerNames []string) *Game {
	fmt.Print("--------------Welcome to GOMonopoly------------------")
	// Erstelle das Spielbrett
	board := NewBoard()

	// Erstelle Spieler
	players := make([]*Player, len(playerNames))
	for i, name := range playerNames {
		players[i] = NewPlayer(name)
	}

	// Mische die Reihenfolge der Spieler
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(players), func(i, j int) {
		players[i], players[j] = players[j], players[i]
	})

	// Setze den Spielzustand
	game := &Game{
		Players:     players,
		Board:       board,
		CurrentTurn: 0,
	}

	return game
}
func GameIsOver(game *Game) bool {
	// Du kannst eine Endbedingung festlegen, z.B. wenn nur noch ein Spieler übrig ist
	remainingPlayers := 0
	for _, player := range game.Players {
		if !player.isBankrupt {
			remainingPlayers++
		}
	}

	return remainingPlayers <= 1
}
