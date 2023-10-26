package monopoly

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Chest struct {
	text string
}

var ChestMap = map[int]*Chest{}

func InitChestCards() {

	content, err := ioutil.ReadFile("chest.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	err = json.Unmarshal(content, &ChestMap)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
}
