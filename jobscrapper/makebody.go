package jobscrapper

import (
	"database/sql"
	"fmt"
	"github.com/msyhu/GobbyIsntFree/aws"
	"github.com/msyhu/GobbyIsntFree/etc"
	_struct "github.com/msyhu/GobbyIsntFree/struct"
	"log"
	"time"
)

func MakeHtmlBody() *string {
	today := time.Now().Format("2006-01-02")

	contents := "<h1>" + today + " 의 개발 채용 일보📰</h1>" +
		"<h2>오늘의 신규 채용</h2><ul>"

	gobbyRdsSecret := aws.GetSecret()

	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true",
		gobbyRdsSecret.User,
		gobbyRdsSecret.Password,
		gobbyRdsSecret.Host,
		gobbyRdsSecret.Database)

	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	etc.CheckErr(err)
	defer db.Close()

	// 오늘 새로 크롤링된 job 조회
	// 오늘 새로 크롤링된 job body 만들어주기

	todayQuery := "SELECT * FROM jobs WHERE START_DATE = '" + today + "'"
	todayRows, err := db.Query(todayQuery)
	etc.CheckErr(err)
	defer todayRows.Close()
	for todayRows.Next() {
		var tempJob _struct.Kakao
		err := todayRows.Scan(&tempJob.Id, &tempJob.Company, &tempJob.Url, &tempJob.EndDate, &tempJob.StartDate, &tempJob.Location, &tempJob.Title)
		if err != nil {
			log.Fatal(err)
		}
		rowHTML := "<li>" +
			"<a href='" + tempJob.Url + "'>" +
			tempJob.Title +
			"</a>" +
			"</li>"
		contents += rowHTML
	}

	// 그외 기존 job 조회
	// 기존 job body 만들어주기
	notTodayQuery := "SELECT * FROM jobs WHERE START_DATE <> '" + today + "'"
	contents += "</ul><h2>기존 채용</h2><ul>"
	beforeRows, err := db.Query(notTodayQuery)
	etc.CheckErr(err)
	defer beforeRows.Close()
	for beforeRows.Next() {
		var tempJob _struct.Kakao
		err := beforeRows.Scan(&tempJob.Id, &tempJob.Company, &tempJob.Url, &tempJob.EndDate, &tempJob.StartDate, &tempJob.Location, &tempJob.Title)
		if err != nil {
			log.Fatal(err)
		}
		rowHTML := "<li>" +
			"<a href='" + tempJob.Url + "'>" +
			tempJob.Title +
			"</a>" +
			"</li>"
		contents += rowHTML
	}
	contents += "</ul>"

	return &contents
}
