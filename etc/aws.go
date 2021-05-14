package etc

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
)

type SecretManager struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Database string `json:"database"`
}

type Subscriber struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func getSecret() *SecretManager {
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
	var gobbyRdsSecret = SecretManager{}
	jsonErr := json.Unmarshal([]byte(secretString), &gobbyRdsSecret)
	CheckErr(jsonErr)

	return &gobbyRdsSecret

}

func GetSubscribers() []Subscriber {
	gobbyRdsSecret := getSecret()

	var connectionString = fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?allowNativePasswords=true",
		gobbyRdsSecret.User,
		gobbyRdsSecret.Password,
		gobbyRdsSecret.Host,
		gobbyRdsSecret.Database)

	// Initialize connection object.
	db, err := sql.Open("mysql", connectionString)
	CheckErr(err)
	defer db.Close()

	err = db.Ping()
	CheckErr(err)
	fmt.Println("Successfully created connection to database.")

	// TODO : 쿼리도 암호화 해야하나?
	rows, err := db.Query("SELECT name, email from subscribers;")
	CheckErr(err)
	defer rows.Close()
	fmt.Println("Reading data:")
	var subscribers []Subscriber

	err = rows.Err()
	CheckErr(err)

	for rows.Next() {
		subscriber := Subscriber{}
		err := rows.Scan(&subscriber.Name, &subscriber.Email)
		CheckErr(err)
		subscribers = append(subscribers, subscriber)
	}

	return subscribers
}
