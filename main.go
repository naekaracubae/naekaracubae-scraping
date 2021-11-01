package main

import (
	"fmt"
	aws2 "github.com/msyhu/naekaracubae-scraping/aws"
	jobscrapper2 "github.com/msyhu/naekaracubae-scraping/jobscrapper"
	"github.com/msyhu/naekaracubae-scraping/struct"
)

type kakaoJob = _struct.Kakao
type lineJob = _struct.Line

func main() {
	jobscrapping()
	//lambda.Start(jobscrapping)
}

// TODO: 회사마다 크롤링, db저장, body 만들기 메서드를 따로 만들어 주었다. 추상화해서 하나로 합칠 수 없을까?
func jobscrapping() string {
	// 1. 크롤링하기
	kakaoJobs, lineJobs := scraping()

	// 2. DB 저장하기
	saveDB(kakaoJobs, lineJobs)

	// 3. MAIL BODY 만들기
	contents := jobscrapper2.MakeHtmlBody()

	// 4. 메일 보내기
	subscribers := aws2.GetSubscribers()
	sendMailResult := aws2.SendMail(contents, subscribers)

	return sendMailResult
}

func scraping() (*[]kakaoJob, *[]lineJob) {
	// 카카오
	kakaoC := make(chan []kakaoJob)
	go jobscrapper2.KakaoCrawling(kakaoC)
	kakaoJobs := <-kakaoC
	fmt.Println(kakaoJobs)
	// 라인
	lineC := make(chan []lineJob)
	go jobscrapper2.LineCrawling(lineC)
	lineJobs := <-lineC
	fmt.Println(lineJobs)

	return &kakaoJobs, &lineJobs
}

func saveDB(kakaoJobs *[]kakaoJob, lineJobs *[]lineJob) {
	// 카카오
	aws2.CheckAndSaveJobForKakao(kakaoJobs)
	// 라인
	aws2.CheckAndSaveJobForLine(lineJobs)
}
