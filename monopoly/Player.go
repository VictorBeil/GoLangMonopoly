package monopoly

import "fmt"

type Player struct {
	money         int
	properties    *[]Property
	currentSquare Property
	isBankrupt    bool
}

func PurchaseProperty() bool {
	return true
}
func SellProperty() bool {
	return true
}
func Move() int {

	return 0
}
func ViewPlayer(player Player) {
	fmt.Printf("%+v\n", player)
}
