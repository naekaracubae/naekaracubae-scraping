package aws_test

import (
	"database/sql"
	"fmt"
	aws2 "github.com/msyhu/naekaracubae-scraping/developerilbo/aws"
	etc2 "github.com/msyhu/naekaracubae-scraping/developerilbo/etc"
	"github.com/msyhu/naekaracubae-scraping/developerilbo/jobscrapper"
	_struct2 "github.com/msyhu/naekaracubae-scraping/developerilbo/struct"
	"log"
	"testing"
)

type lineJob = _struct2.Line

var testLineStruct = lineJob{
	Title:    "Global E-Commerce 경력채용",
	EndDate:  "채용시까지",
	Location: "Bundang",
	Company:  "LINE PlusDesign",
	Url:      "https://careers.linecorp.com/ko/jobs/862",
	Id:       "862",
}

func Test_IsJobExistForLine(t *testing.T) {

	gobbyRdsSecret := aws2.GetSecret()

	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true",
		gobbyRdsSecret.User,
		gobbyRdsSecret.Password,
		gobbyRdsSecret.Host,
		gobbyRdsSecret.Database)

	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	etc2.CheckErr(err)
	defer db.Close()

	result := aws2.IsJobExistForLine(&testLineStruct, db)

	if result != true {
		t.Error("Wrong result")
	}
}

func Test_SaveJobForLine(t *testing.T) {
	gobbyRdsSecret := aws2.GetSecret()

	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true",
		gobbyRdsSecret.User,
		gobbyRdsSecret.Password,
		gobbyRdsSecret.Host,
		gobbyRdsSecret.Database)

	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	etc2.CheckErr(err)
	defer db.Close()

	result := aws2.SaveJobForLine(&testLineStruct, db)

	if result != true {
		t.Error("Wrong result")
	}
}

func Test_CheckAndSaveJobForLine(t *testing.T) {

	lineC := make(chan []lineJob)
	go jobscrapper.LineCrawling(lineC)
	lineJobs := <-lineC
	log.Println(lineJobs)

	aws2.CheckAndSaveJobForLine(&lineJobs)

}
