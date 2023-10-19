package monopoly

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Property struct {
	Owner        int	`json:"owner"`
	Name         string	`json:"name"`
	Type         string	`json:"type"`
	Cost         int	`json:"cost"`
	Color        string	`json:"color"`
	Rent 		 int[]	`json:"rent"`
	Group		 int[]	`json:"group"`
	CostHouse	 int	`json:"house"`
}

var PropertyMap = map[int]*Property{}

func InitProperties() {

	content, err := ioutil.ReadFile("property.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	err = json.Unmarshal(content, &PropertyMap)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
}
