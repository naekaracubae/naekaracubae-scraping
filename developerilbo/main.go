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
	// 1. 크롤링하기
	kakaoJobs, lineJobs := scraping()

	// 2. DB 저장하기
	saveDB(kakaoJobs, lineJobs)

	// 3. MAIL BODY 만들기
	contents := jobscrapper.MakeHtmlBody()

	// 4. 메일 보내기
	sendMailResult := aws.SendMail(contents)

	return sendMailResult
}

func scraping() (*[]kakaoJob, *[]lineJob) {
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

	return &kakaoJobs, &lineJobs
}

func saveDB(kakaoJobs *[]kakaoJob, lineJobs *[]lineJob) {
	// 카카오
	aws.CheckAndSaveJobForKakao(kakaoJobs)
	// 라인
	aws.CheckAndSaveJobForLine(lineJobs)
}
