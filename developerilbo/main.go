package main

import (
	"fmt"
	aws "github.com/msyhu/naekaracubae-scraping/developerilbo/aws"
	jobscrapper "github.com/msyhu/naekaracubae-scraping/developerilbo/jobscrapper"
	_struct "github.com/msyhu/naekaracubae-scraping/developerilbo/struct"
)

type kakaoJob = _struct.Kakao
type lineJob = _struct.Line

func main() {
	jobscrapping()
}

// TODO: 회사마다 크롤링, db저장, body 만들기 메서드를 따로 만들어 주었다. 추상화해서 하나로 합칠 수 없을까?
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
	// 카카오
	aws.CheckAndSaveJobForKakao(&kakaoJobs)
	// 라인
	aws.CheckAndSaveJobForLine(&lineJobs)

	contents := jobscrapper.MakeHtmlBody()

	subscribers := aws.GetSubscribers()
	sendMailResult := aws.SendMail(contents, subscribers)

	return sendMailResult
}
