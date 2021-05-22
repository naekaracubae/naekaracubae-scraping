package aws

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	_ "github.com/go-sql-driver/mysql"
	"github.com/msyhu/GobbyIsntFree/etc"
	_struct "github.com/msyhu/GobbyIsntFree/struct"
	"log"
	"time"
)

type kakaoExtractedJob = _struct.Kakao

func GetSecret() *_struct.SecretManager {
	secretName := "GOBBY_RDS_SECRETS"
	region := "ap-northeast-2"

	//Create a Secrets Manager client
	svc := secretsmanager.New(session.New(),
		aws.NewConfig().WithRegion(region))
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	// In this sample we only handle the specific exceptions for the 'GetSecretValue' API.
	// See https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html

	result, err := svc.GetSecretValue(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case secretsmanager.ErrCodeDecryptionFailure:
				// Secrets Manager can't decrypt the protected secret text using the provided KMS key.
				fmt.Println(secretsmanager.ErrCodeDecryptionFailure, aerr.Error())

			case secretsmanager.ErrCodeInternalServiceError:
				// An error occurred on the server side.
				fmt.Println(secretsmanager.ErrCodeInternalServiceError, aerr.Error())

			case secretsmanager.ErrCodeInvalidParameterException:
				// You provided an invalid value for a parameter.
				fmt.Println(secretsmanager.ErrCodeInvalidParameterException, aerr.Error())

			case secretsmanager.ErrCodeInvalidRequestException:
				// You provided a parameter value that is not valid for the current state of the resource.
				fmt.Println(secretsmanager.ErrCodeInvalidRequestException, aerr.Error())

			case secretsmanager.ErrCodeResourceNotFoundException:
				// We can't find the resource that you asked for.
				fmt.Println(secretsmanager.ErrCodeResourceNotFoundException, aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}

	}

	// Decrypts secret using the associated KMS CMK.
	// Depending on whether the secret is a string or binary, one of these fields will be populated.
	var secretString, _ string
	if result.SecretString != nil {
		secretString = *result.SecretString
	} else {
		decodedBinarySecretBytes := make([]byte, base64.StdEncoding.DecodedLen(len(result.SecretBinary)))
		len, err := base64.StdEncoding.Decode(decodedBinarySecretBytes, result.SecretBinary)
		if err != nil {
			fmt.Println("Base64 Decode Error:", err)
			return nil
		}
		_ = string(decodedBinarySecretBytes[:len])
	}

	// Your code goes here.
	var gobbyRdsSecret = _struct.SecretManager{}
	jsonErr := json.Unmarshal([]byte(secretString), &gobbyRdsSecret)
	etc.CheckErr(jsonErr)

	return &gobbyRdsSecret

}

func GetSubscribers() []_struct.Subscriber {
	gobbyRdsSecret := GetSecret()

	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true",
		gobbyRdsSecret.User,
		gobbyRdsSecret.Password,
		gobbyRdsSecret.Host,
		gobbyRdsSecret.Database)

	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	etc.CheckErr(err)
	defer db.Close()

	// query := "SELECT name, email from subscribers;"
	query := "SELECT name, email from test_subscribers;"
	rows, err := db.Query(query)
	etc.CheckErr(err)
	defer rows.Close()
	fmt.Println("Reading data:")
	var subscribers []_struct.Subscriber

	err = rows.Err()
	etc.CheckErr(err)

	for rows.Next() {
		subscriber := _struct.Subscriber{}
		err := rows.Scan(&subscriber.Name, &subscriber.Email)
		etc.CheckErr(err)
		subscribers = append(subscribers, subscriber)
	}

	return subscribers
}

func CheckAndSaveJob(kakaoJobs *[]kakaoExtractedJob) {

	gobbyRdsSecret := GetSecret()

	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true",
		gobbyRdsSecret.User,
		gobbyRdsSecret.Password,
		gobbyRdsSecret.Host,
		gobbyRdsSecret.Database)

	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	etc.CheckErr(err)
	defer db.Close()

	// 테이블을 ID로 조회해서 없는 경우 DB에 새로 저장한다.
	for _, kakaoJob := range *kakaoJobs {
		if !IsJobExist(&kakaoJob, db) {
			SaveJob(&kakaoJob, db)
		}
	}

}

func IsJobExist(kakaoJobs *kakaoExtractedJob, db *sql.DB) bool {

	id := kakaoJobs.Id

	query := "SELECT id FROM jobs WHERE ID='" + id + "'"
	var row string
	queryErr := db.QueryRow(query).Scan(&row)
	if queryErr != nil {
		return false
	}

	return true
}

func SaveJob(kakaoJobs *kakaoExtractedJob, db *sql.DB) bool {
	startDate := time.Now().Format("2006-01-02")
	result, err := db.Exec("INSERT INTO jobs VALUES (?, ?, ?, ?, ?, ?, ?)", kakaoJobs.Id, kakaoJobs.Company, kakaoJobs.Url, kakaoJobs.EndDate, startDate, kakaoJobs.Location, kakaoJobs.Title)
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
