package main

import (
	"fmt"
	aws "github.com/msyhu/GobbyIsntFree/developerilbo/aws"
	jobscrapper "github.com/msyhu/GobbyIsntFree/developerilbo/jobscrapper"
	_struct "github.com/msyhu/GobbyIsntFree/developerilbo/struct"
)

type kakaoJob = _struct.Kakao
type lineJob = _struct.Line

func main() {
	jobscrapping()
}

func jobscrapping() string {
	// 크롤링하기
	kakaoC := make(chan []kakaoJob)
	go jobscrapper.KakaoCrawling(kakaoC)
	kakaoJobs := <-kakaoC

	fmt.Println(kakaoJobs)

	// DB 저장하기
	aws.CheckAndSaveKakaoJobs(&kakaoJobs)

	contents := jobscrapper.MakeHtmlBody()

	// 메일 보내기 : 함수 하나로 만들것
	subscribers := aws.GetSubscribers()
	sendMailResult := aws.SendMail(contents, subscribers)

	return sendMailResult
}
