package aws_test

import (
	"database/sql"
	"fmt"
	aws2 "github.com/msyhu/GobbyIsntFree/developerilbo/aws"
	etc2 "github.com/msyhu/GobbyIsntFree/developerilbo/etc"
	_struct2 "github.com/msyhu/GobbyIsntFree/developerilbo/struct"
	"testing"
)

type kakaoJob = _struct2.Kakao

func TestGetSubscribers(t *testing.T) {
	subscribers := aws2.GetSubscribers()
	fmt.Println(subscribers)
}

var testKakaoStruct = kakaoJob{
	Title:    "test",
	EndDate:  "채용시까지",
	Location: "판교",
	Company:  "kakao",
	Url:      "https://careers.kakao.com/jobs/P-9349?part=TECHNOLOGY&company=ALL",
	JobType:  "정규직",
	Id:       "P-9349",
}

func TestIsJobExist(t *testing.T) {

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

	result := aws2.IsJobExist(&testKakaoStruct, db)

	if result != true {
		t.Error("Wrong result")
	}
}

func TestSaveJob(t *testing.T) {
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

	result := aws2.SaveJob(&testKakaoStruct, db)

	if result != true {
		t.Error("Wrong result")
	}
}

func TestGetSecret(t *testing.T) {
	gobbyRdsSecret := aws2.GetSecret()

	if gobbyRdsSecret == nil {
		t.Error("Wrong result")
	}
}
