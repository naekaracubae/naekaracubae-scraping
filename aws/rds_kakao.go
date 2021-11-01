package aws

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	etc2 "github.com/msyhu/naekaracubae-scraping/etc"
	_struct2 "github.com/msyhu/naekaracubae-scraping/struct"
	"log"
	"time"
)

type kakaoExtractedJob = _struct2.Kakao

func CheckAndSaveJobForKakao(kakaoJobs *[]kakaoExtractedJob) {

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
	for _, kakaoJob := range *kakaoJobs {
		if !IsJobExistForKakao(&kakaoJob, db) {
			SaveJobForKakao(&kakaoJob, db)
		} else {
			// 이미 존재하면, 마지막 있었던 날짜(LAST_EXIST_DATE) 최신화 시켜주기.
			// 메일 보낼때, LAST_EXIST_DATE가 오늘 날짜인 ROW 만 전송한다.
			updateLastExistDateForKakao(&kakaoJob, db)
		}
	}

}

func updateLastExistDateForKakao(kakaoJob *kakaoExtractedJob, db *sql.DB) bool {
	today := time.Now().Format("2006-01-02")
	result, err := db.Exec("UPDATE jobs SET LAST_EXIST_DATE=? WHERE ID=?", today, kakaoJob.Id)
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

func IsJobExistForKakao(kakaoJobs *kakaoExtractedJob, db *sql.DB) bool {

	id := kakaoJobs.Id

	query := "SELECT id FROM jobs WHERE ID='" + id + "'"
	var row string
	queryErr := db.QueryRow(query).Scan(&row)
	if queryErr != nil {
		return false
	}

	return true
}

func SaveJobForKakao(kakaoJobs *kakaoExtractedJob, db *sql.DB) bool {
	today := time.Now().Format("2006-01-02")
	result, err := db.Exec("INSERT INTO jobs VALUES (?, ?, ?, ?, ?, ?, ?, ?)", kakaoJobs.Id, kakaoJobs.Company, kakaoJobs.Url, kakaoJobs.EndDate, today, kakaoJobs.Location, kakaoJobs.Title, today)
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
