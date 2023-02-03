package jobs

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func CreateJob(username, password, host, sqlCommand, token string, limit int) (string, error) {
	// Define the request body
	body := `{"commands": "` + sqlCommand + `", "limit": ` + strconv.Itoa(limit) + `, "separator": ";", "stop_on_error": "no"}`
	jsonStr := []byte(body)
	req, err := http.NewRequest("POST", "https://"+host+"/dbapi/v4/sql_jobs", bytes.NewBuffer(jsonStr))
	if err != nil {
		return "error: ", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

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
	return data["id"].(string), nil
}
