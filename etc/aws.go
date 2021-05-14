package etc

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rdsdataservice"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func MakeConnection() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-2"),
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	SQLStatement := `SELECT * FROM subscribers;`

	rdsdataservice_client := rdsdataservice.New(sess)
	req, resp := rdsdataservice_client.ExecuteStatementRequest(&rdsdataservice.ExecuteStatementInput{
		Database:    aws.String("gobbyisntfree"),
		ResourceArn: aws.String("arn:aws:rds:ap-northeast-2:685320160057:db:gobbyisntfree"),
		//SecretArn:   aws.String("arn:aws:secretsmanager:us-east-2:9xxxxxxxx9:secret:RDS_Credentials-IZOXv0"),
		Sql: aws.String(SQLStatement),
	})

	err1 := req.Send()
	if err1 == nil { // resp is now filled
		fmt.Println("Response:", resp)
	} else {
		fmt.Println("error:", err1)
	}
}

func GetSubscribers() {

}

func GetSenders() {

}
