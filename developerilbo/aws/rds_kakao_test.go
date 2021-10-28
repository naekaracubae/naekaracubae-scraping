package aws_test

import (
	"database/sql"
	"fmt"
	aws2 "github.com/msyhu/naekaracubae-scraping/developerilbo/aws"
	etc2 "github.com/msyhu/naekaracubae-scraping/developerilbo/etc"
	_struct2 "github.com/msyhu/naekaracubae-scraping/developerilbo/struct"
	"testing"
)

type kakaoJob = _struct2.Kakao

var testKakaoStruct = kakaoJob{
	Title:    "test",
	EndDate:  "채용시까지",
	Location: "판교",
	Company:  "kakao",
	Url:      "https://careers.kakao.com/jobs/P-9349?part=TECHNOLOGY&company=ALL",
	JobType:  "정규직",
	Id:       "P-9349",
}

func Test_IsJobExistForKakao(t *testing.T) {

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

	result := aws2.IsJobExistForKakao(&testKakaoStruct, db)

	if result != true {
		t.Error("Wrong result")
	}
}

func Test_SaveJobForKakao(t *testing.T) {
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

	result := aws2.SaveJobForKakao(&testKakaoStruct, db)

	if result != true {
		t.Error("Wrong result")
	}
}
