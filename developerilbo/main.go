package main

import (
	"fmt"
	aws2 "github.com/msyhu/naekaracubae-scraping/developerilbo/aws"
	jobscrapper "github.com/msyhu/naekaracubae-scraping/developerilbo/jobscrapper"
	_struct "github.com/msyhu/naekaracubae-scraping/developerilbo/struct"
)

type kakaoJob = _struct.Kakao
type lineJob = _struct.Line

func main() {
	jobscrapping()
}

func jobscrapping() string {
	// 크롤링하기
	// 카카오
	kakaoC := make(chan []kakaoJob)
	go jobscrapper.KakaoCrawling(kakaoC)
	kakaoJobs := <-kakaoC
	fmt.Println(kakaoJobs)
	// 라인
	lineC := make(chan []lineJob)
	go jobscrapper.LineCrawling(lineC)
	lineJobs := <-lineC
	fmt.Println(lineJobs)

	// DB 저장하기
	aws2.CheckAndSaveJob(&kakaoJobs)

	contents := jobscrapper.MakeHtmlBody()

	// 메일 보내기 : 함수 하나로 만들것
	subscribers := aws2.GetSubscribers()
	sendMailResult := aws2.SendMail(contents, subscribers)

	return sendMailResult
}
