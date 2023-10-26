package monopoly

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Chance struct {
	text string
}

var ChanceMap = map[int]*Chance{}

func InitChanceCards() {

	content, err := ioutil.ReadFile("chance.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	err = json.Unmarshal(content, &ChanceMap)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
}
