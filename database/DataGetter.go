package data

import (
	"os"
	"io/ioutil"
	"../model"
	"fmt"
	"encoding/json"
)

var JSON []byte
var PLACES []model.Place

func GetData() {
	// load options data from json file
	file, e := ioutil.ReadFile("./database/data.json")

	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	JSON = file

	// load places data from json file
	file, er := ioutil.ReadFile("./database/places.json")

	if er != nil {
		fmt.Printf("File error: %v\n", er)
		os.Exit(1)
	}

	json.Unmarshal(file, &PLACES)
}