package game

var m_playerList [8]player
var m_numOfPlayers int
var m_numOfCurrentPlayer int
var m_communityDeck *Deck
var m_chanceDeck *Deck
var m_board *board
var m_dice *dice

func CheckWin() bool {
	return true
}
func GetNumOfPlayers() int {
	return 0
}
func GetCurrentPlayer() *player {

	currentPlayer := &player{
		name:  "John Doe",
		score: 100,
	}
	return currentPlayer
}
func GetPlayerNumber(number int) *player {

	currentPlayer := &player{
		name:  "John Doe",
		score: 100,
	}
	return currentPlayer
}
func RollDice() {
	m_dice += random.
}
