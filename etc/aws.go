package etc

import (
	"database/sql"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/service/rds/rdsutils"

	_ "github.com/go-sql-driver/mysql"
)

func MakeConnection() {
	dbName := "gobbyisntfree"
	dbUser := "msyhu"
	dbHost := "gobbyisntfree.ccttcm80dlu1.ap-northeast-2.rds.amazonaws.com"
	dbPort := 3306
	dbEndpoint := fmt.Sprintf("%s:%d", dbHost, dbPort)
	region := "ap-northeast-2"

	creds := credentials.NewEnvCredentials()
	//credValue, err := creds.Get()
	//if err != nil {
	//	panic(err.Error())
	//}

	//fmt.Printf("%+v", credValue)

	authToken, err := rdsutils.BuildAuthToken(dbEndpoint, region, dbUser, creds)
	if err != nil {
		panic(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=true&allowCleartextPasswords=true",
		dbUser, authToken, dbEndpoint, dbName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func GetSubscribers() {

}

func GetSenders() {

}
