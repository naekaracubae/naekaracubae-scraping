package jobscrapper

import (
	"database/sql"
	"fmt"
	aws2 "github.com/msyhu/naekaracubae-scraping/developerilbo/aws"
	etc2 "github.com/msyhu/naekaracubae-scraping/developerilbo/etc"
	_struct2 "github.com/msyhu/naekaracubae-scraping/developerilbo/struct"
	"log"
	"time"
)

func MakeHtmlBody() *string {
	today := time.Now().Format("2006-01-02")

	contents := "<h1>" + "[네,카라쿠배] " + today + " 개발자 채용 일보📰</h1>" +
		"<h2>오늘의 신규 채용</h2>"

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

	// 오늘 새로 크롤링된 job
	// 카카오
	contents += "<h3>카카오</h3><ul>"
	// 오늘 새로 크롤링된 job body 만들어주기
	todayQuery := "SELECT * FROM jobs WHERE START_DATE = '" + today + "'"
	todayRows, err := db.Query(todayQuery)
	etc2.CheckErr(err)
	defer todayRows.Close()
	for todayRows.Next() {
		var tempJob _struct2.Kakao
		err := todayRows.Scan(&tempJob.Id, &tempJob.Company, &tempJob.Url, &tempJob.EndDate, &tempJob.StartDate, &tempJob.Location, &tempJob.Title, &tempJob.LastExistDate)
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

	// 라인
	contents += "<h3>라인</h3><ul>"
	// 오늘 새로 크롤링된 job body 만들어주기
	todayQueryForLine := "SELECT * FROM LINE WHERE START_DATE = '" + today + "'"
	todayRowsForLine, err := db.Query(todayQueryForLine)
	etc2.CheckErr(err)
	defer todayRowsForLine.Close()
	for todayRowsForLine.Next() {
		var tempJob _struct2.Line
		err := todayRowsForLine.Scan(&tempJob.Id, &tempJob.Company, &tempJob.Url, &tempJob.EndDate, &tempJob.StartDate, &tempJob.Location, &tempJob.Title, &tempJob.LastExistDate)
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

	// 그외 기존 job 조회
	// 카카오
	// 기존 job body 만들어주기
	notTodayQuery := "SELECT * FROM jobs WHERE START_DATE <> '" + today + "' AND LAST_EXIST_DATE = '" + today + "'"
	contents += "<h2>기존 채용</h2>"
	contents += "<h3>카카오</h3><ul>"
	beforeRows, err := db.Query(notTodayQuery)
	etc2.CheckErr(err)
	defer beforeRows.Close()
	for beforeRows.Next() {
		var tempJob _struct2.Kakao
		err := beforeRows.Scan(&tempJob.Id, &tempJob.Company, &tempJob.Url, &tempJob.EndDate, &tempJob.StartDate, &tempJob.Location, &tempJob.Title, &tempJob.LastExistDate)
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
