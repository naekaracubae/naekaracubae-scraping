package main

import (
	"fmt"
	"github.com/msyhu/GobbyIsntFree/aws"
	"github.com/msyhu/GobbyIsntFree/etc"
	"github.com/msyhu/GobbyIsntFree/jobscrapper"
	_struct "github.com/msyhu/GobbyIsntFree/struct"
)

type kakaoExtractedJob = _struct.Kakao

func main() {
	jobscrapping()
}

func jobscrapping() {
	// 크롤링하기
	kakaoC := make(chan []kakaoExtractedJob)
	go jobscrapper.KakaoCrawling(kakaoC)
	kakaoJobs := <-kakaoC

	// DB 저장하기
	aws.CheckAndSaveJob(&kakaoJobs)

	// 메일 body 가공하기 : 수정요망
	contents := etc.StructToStr(&kakaoJobs)
	fmt.Println(contents)

	// 메일 보내기 : 함수 하나로 만들것
	subscribers := aws.GetSubscribers()
	fmt.Println(subscribers)
	aws.SendMail(contents, subscribers)
}
