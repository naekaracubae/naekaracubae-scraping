package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	aws2 "github.com/msyhu/GobbyIsntFree/developerilbo/aws"
	jobscrapper2 "github.com/msyhu/GobbyIsntFree/developerilbo/jobscrapper"
	_struct2 "github.com/msyhu/GobbyIsntFree/developerilbo/struct"
)

type kakaoExtractedJob = _struct2.Kakao
type lineExtractedJob = _struct2.Line

func main() {
	//lambda.Start(Handler)
	//KakaoScrapping()
	LineScrapping()
	MakeAndSendMail()
}

// Response is of type APIGatewayProxyResponse since we're leveraging the
// AWS Lambda Proxy Request functionality (default behavior)
//
// https://serverless.com/framework/docs/providers/aws/events/apigateway/#lambda-proxy-integration
type Response events.APIGatewayProxyResponse

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {
	var buf bytes.Buffer

	go KakaoScrapping()

	body, err := json.Marshal(map[string]interface{}{
		"message": "ok",
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func MakeAndSendMail() string {
	contents := jobscrapper2.MakeHtmlBody()

	// 메일 보내기 : 함수 하나로 만들것
	subscribers := aws2.GetSubscribers()
	sendMailResult := aws2.SendMail(contents, subscribers)

	return sendMailResult
}

func KakaoScrapping() {
	// 크롤링하기
	kakaoC := make(chan []kakaoExtractedJob)
	go jobscrapper2.KakaoCrawling(kakaoC)
	kakaoJobs := <-kakaoC

	fmt.Println(kakaoJobs)

	// DB 저장하기
	aws2.CheckAndSaveJob(&kakaoJobs)

}

func LineScrapping() {
	// 크롤링하기
	lineC := make(chan []lineExtractedJob)
	go jobscrapper2.LineCrawling(lineC)
	lineJobs := <-lineC

	fmt.Println(lineJobs)

	// DB 저장하기
	// TODO : 인터페이스로 kakao 구조체와 line 구조체 동시에 받을수 있도록 조상 생성
	//aws2.CheckAndSaveJob(&lineJobs)

}
