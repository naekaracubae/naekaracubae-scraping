package aws_test

import (
	"database/sql"
	"fmt"
	"github.com/msyhu/GobbyIsntFree/aws"
	"github.com/msyhu/GobbyIsntFree/etc"
	_struct "github.com/msyhu/GobbyIsntFree/struct"
	"testing"
)

type kakaoJob = _struct.Kakao

func TestGetSubscribers(t *testing.T) {
	aws.GetSubscribers()
}

func TestIsJobExist(t *testing.T) {

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

	testKakaoStruct := kakaoJob{
		Title:    "test",
		EndDate:  "채용시까지",
		Location: "판교",
		Company:  "kakao",
		Url:      "https://careers.kakao.com/jobs/P-9349?part=TECHNOLOGY&company=ALL",
		JobType:  "정규직",
		Id:       "P-9349",
	}

	result := aws.IsJobExist(&testKakaoStruct, db)

	if result != true {
		t.Error("Wrong result")
	}
}
