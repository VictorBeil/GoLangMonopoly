package game

import (
	"math/rand"
	"strconv"
)

var m_dices [2]int

func roll() {
	// Generiere eine zufällige Zahl zwischen 1 und 6 (einschließlich)
	m_dices[0] = rand.Intn(6-1) + 1
	m_dices[1] = rand.Intn(6-1) + 1
}
func checkDouble() bool {
	if m_dices[0] == m_dices[1] {
		return true
	}
	return false
}
func getTotal() int {
	return m_dices[0] + m_dices[1]
}
func getLastRollString() string {
	total := getTotal()

	// Konvertiere die Zahl in einen String
	numbersAsString := strconv.Itoa(total)

	return numbersAsString
}
