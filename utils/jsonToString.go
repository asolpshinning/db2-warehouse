package utils

import (
	"encoding/json"
)

func JSONObjectToString(obj map[string]interface{}) (string, error) {
	jsonBytes, err := json.MarshalIndent(obj, "", "    ")
	if err != nil {
		return "", err
	}
	jsonStr := string(jsonBytes)
	//loop through jsonStr, and replace `"` with `\"` and return a new string that has `\"` instead of `"`
	for i := 0; i < len(jsonStr); i++ {
		if jsonStr[i] == '"' {
			jsonStr = jsonStr[:i] + `\"` + jsonStr[i+1:]
			i++
		}
	}
	return jsonStr, nil
}
