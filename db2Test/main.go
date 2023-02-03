package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/asolpshinning/db2-warehouse/utils"

	"github.com/asolpshinning/db2-warehouse/jobs"
)

func main() {

	userName := utils.GoDotEnv("username")
	password := utils.GoDotEnv("password")
	host := utils.GoDotEnv("host")
	//sqlCommand := `select APPLICATION_NAME, IMAP_NUMBER, APP_DOMAIN FROM COE.APPLICATION_DIM fetch first 5 rows only;select TECHNICAL_OWNER, PRODUCT_OWNER, LOAD_DATE from COE.APPLICATION_DIM fetch first 5 rows only;`
	//sqlCommand := `SELECT * FROM COE.CIRRUS_AUTO_SCALING_V`
	sqlCommand := `SELECT * FROM COE.CIRRUS_AUTO_SCALING_V WHERE ACTIVE = 'Y' AND ENV = 'prod' OFFSET 1000 ROWS FETCH NEXT 100 ROWS ONLY`
	limit := 100

	result, err := jobs.GetResultFromJob(userName, password, host, sqlCommand, limit)
	if err != nil {
		panic(err)
	}
	//convert the result to a json and print it
	jsonResult, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("result.json", jsonResult, 0644)
	fmt.Println("Done! Check the result.json file")
}
