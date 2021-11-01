package aws_test

import (
	"database/sql"
	"fmt"
	"github.com/msyhu/naekaracubae-scraping/aws"
	etc2 "github.com/msyhu/naekaracubae-scraping/etc"
	"github.com/msyhu/naekaracubae-scraping/jobscrapper"
	_struct2 "github.com/msyhu/naekaracubae-scraping/struct"
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

	gobbyRdsSecret := aws.GetSecret()

	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true",
		gobbyRdsSecret.User,
		gobbyRdsSecret.Password,
		gobbyRdsSecret.Host,
		gobbyRdsSecret.Database)

	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	etc2.CheckErr(err)
	defer db.Close()

	result := aws.IsJobExistForLine(&testLineStruct, db)

	if result != true {
		t.Error("Wrong result")
	}
}

func Test_SaveJobForLine(t *testing.T) {
	gobbyRdsSecret := aws.GetSecret()

	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true",
		gobbyRdsSecret.User,
		gobbyRdsSecret.Password,
		gobbyRdsSecret.Host,
		gobbyRdsSecret.Database)

	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	etc2.CheckErr(err)
	defer db.Close()

	result := aws.SaveJobForLine(&testLineStruct, db)

	if result != true {
		t.Error("Wrong result")
	}
}

// TODO: 실제 db에 집어넣어 버리는데, 테스트 db 적용하고 테스트 후에는 데이터 삭제하는 코드추가
func Test_CheckAndSaveJobForLine(t *testing.T) {

	lineC := make(chan []lineJob)
	go jobscrapper.LineCrawling(lineC)
	lineJobs := <-lineC
	log.Println(lineJobs)

	aws.CheckAndSaveJobForLine(&lineJobs)

}
