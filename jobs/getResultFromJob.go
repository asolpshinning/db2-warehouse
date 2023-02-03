package jobs

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/asolpshinning/db2-warehouse/jobs/auth"
)

func GetResultFromJob(username, password, host, sqlCommand string, limit int) (map[string]interface{}, error) {

	token, err := auth.RequestToken(username, password, host)
	if err != nil {
		return nil, err
	}

	id, err := CreateJob(username, password, host, sqlCommand, token, limit)
	if err != nil {
		return nil, err
	}

	// Define the request body
	req, err := http.NewRequest("GET", "https://"+host+"/dbapi/v4/sql_jobs/"+id, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Print the response
	respBody, _ := ioutil.ReadAll(resp.Body)

	//convert to map[string]interface{}
	var data map[string]interface{}
	err = json.Unmarshal(respBody, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
