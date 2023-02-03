package auth

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func RequestToken(username, password, host string) (string, error) {
	// Define the request body
	body := `{ "userid": "` + username + `", "password": "` + password + `"}`
	jsonStr := []byte(body)
	req, err := http.NewRequest("POST", "https://"+host+"/dbapi/v4/auth/tokens", bytes.NewBuffer(jsonStr))
	if err != nil {
		return "error: ", err
	}
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "error: ", err
	}
	defer resp.Body.Close()

	// Print the response
	respBody, _ := ioutil.ReadAll(resp.Body)
	var data map[string]interface{}
	if err := json.Unmarshal(respBody, &data); err != nil {
		return "error: ", err
	}
	return data["token"].(string), nil
}
