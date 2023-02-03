package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func ReadTokenValue(jsonFileName string) string {
	// Read the token from the json file
	jsonFile, err := ioutil.ReadFile(jsonFileName)
	if err != nil {
		fmt.Println(err)
	}
	var data map[string]interface{}
	if err := json.Unmarshal(jsonFile, &data); err != nil {
		fmt.Println(err)
	}
	return data["token"].(string)
}
