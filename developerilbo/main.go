package main

import (
	"fmt"
	aws2 "github.com/msyhu/GobbyIsntFree/developerilbo/aws"
	jobscrapper2 "github.com/msyhu/GobbyIsntFree/developerilbo/jobscrapper"
	_struct2 "github.com/msyhu/GobbyIsntFree/developerilbo/struct"
)

type kakaoExtractedJob = _struct2.Kakao

func main() {
	jobscrapping()
}

func jobscrapping() string {
	// 크롤링하기
	kakaoC := make(chan []kakaoExtractedJob)
	go jobscrapper2.KakaoCrawling(kakaoC)
	kakaoJobs := <-kakaoC

	fmt.Println(kakaoJobs)

	// DB 저장하기
	aws2.CheckAndSaveJob(&kakaoJobs)

	contents := jobscrapper2.MakeHtmlBody()

	// 메일 보내기 : 함수 하나로 만들것
	subscribers := aws2.GetSubscribers()
	sendMailResult := aws2.SendMail(contents, subscribers)

	return sendMailResult
}
