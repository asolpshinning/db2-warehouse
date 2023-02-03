package main

import (
	"encoding/json"
	"fmt"

	"github.com/asolpshinning/db2-warehouse/utils"

	"github.com/asolpshinning/db2-warehouse/jobs"
)

func main() {

	userName := utils.GoDotEnv("username")
	password := utils.GoDotEnv("password")
	host := utils.GoDotEnv("host")
	sqlCommand := `select APPLICATION_NAME, IMAP_NUMBER, APP_DOMAIN FROM COE.APPLICATION_DIM fetch first 5 rows only;select TECHNICAL_OWNER, PRODUCT_OWNER, LOAD_DATE from COE.APPLICATION_DIM fetch first 5 rows only;`
	/* sqlCommand2 := `SELECT
		cad.CIRRUS_APP_NAME ,
		cad.ENV ,
		cpd.NAMESPACE ,
		cad.PROJECT_ID ,
		cad.QUOTA ,
		cad.HS_ACTIVE ,
		cad.IMAP_NUMBER ,
		ad.APP_DOMAIN ,
		ad.APP_DOMAIN_1 ,
		ad.APP_DOMAIN_2 ,
		ad.PRODUCT_OWNER ,
		ad.TECHNICAL_OWNER ,
		cad.START_DATE ,
		cad.END_DATE,
		cad.ACTIVE
	FROM
		COE.CIRRUS_APPLICATION_DIM cad
	JOIN COE.CIRRUS_PROJECT_DIM cpd
	ON
		cpd.PROJECT_ID = cad.PROJECT_ID
	JOIN COE.APPLICATION_DIM ad
	ON
		ad.IMAP_NUMBER = cpd.IMAP_NUMBER` */

	limit := 10

	result, err := jobs.GetResultFromJob(userName, password, host, sqlCommand, limit)
	if err != nil {
		panic(err)
	}
	//convert the result to a json and print it
	jsonResult, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonResult))

}
