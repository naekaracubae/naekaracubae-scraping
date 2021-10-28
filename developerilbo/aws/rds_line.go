package aws

import (
	"database/sql"
	"log"
	"time"
)

func CheckAndSaveJobForLine(lineJobs *[]lineExtractedJob) {

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
