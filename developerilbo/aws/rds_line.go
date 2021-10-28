package aws

import (
	"database/sql"
	"fmt"
	etc2 "github.com/msyhu/naekaracubae-scraping/developerilbo/etc"
	_struct2 "github.com/msyhu/naekaracubae-scraping/developerilbo/struct"
	"log"
	"time"
)

type lineExtractedJob = _struct2.Line

func CheckAndSaveJobForLine(lineJobs *[]lineExtractedJob) {

	gobbyRdsSecret := GetSecret()

	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true",
		gobbyRdsSecret.User,
		gobbyRdsSecret.Password,
		gobbyRdsSecret.Host,
		gobbyRdsSecret.Database)

	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	etc2.CheckErr(err)
	defer db.Close()

	// 테이블을 ID로 조회해서 없는 경우 DB에 새로 저장한다.
	for _, lineJob := range *lineJobs {
		if !IsJobExistForLine(&lineJob, db) {
			SaveJobForLine(&lineJob, db)
		} else {
			// 이미 존재하면, 마지막 있었던 날짜(LAST_EXIST_DATE) 최신화 시켜주기.
			// 메일 보낼때, LAST_EXIST_DATE가 오늘 날짜인 ROW 만 전송한다.
			updateLastExistDateForLine(&lineJob, db)
		}
	}
}

func SaveJobForLine(lineJobs *lineExtractedJob, db *sql.DB) bool {
	today := time.Now().Format("2006-01-02")
	result, err := db.Exec("INSERT INTO LINE VALUES (?, ?, ?, ?, ?, ?, ?, ?)", lineJobs.Id, lineJobs.Company, lineJobs.Url, lineJobs.EndDate, today, lineJobs.Location, lineJobs.Title, today)
	if err != nil {
		log.Fatal(err)
	}

	// sql.Result.RowsAffected() 체크
	n, err := result.RowsAffected()
	if n == 1 {
		return true
	}

	return false
}

func IsJobExistForLine(lineJobs *lineExtractedJob, db *sql.DB) bool {

	id := lineJobs.Id

	query := "SELECT id FROM LINE WHERE ID='" + id + "'"
	var row string
	queryErr := db.QueryRow(query).Scan(&row)
	if queryErr != nil {
		return false
	}

	return true
}

func updateLastExistDateForLine(lineJob *lineExtractedJob, db *sql.DB) bool {
	today := time.Now().Format("2006-01-02")
	result, err := db.Exec("UPDATE LINE SET LAST_EXIST_DATE=? WHERE ID=?", today, lineJob.Id)
	if err != nil {
		log.Fatal(err)
	}

	// sql.Result.RowsAffected() 체크
	n, err := result.RowsAffected()
	if n == 1 {
		return true
	}

	return false
}
