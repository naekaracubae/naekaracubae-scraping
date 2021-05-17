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
	kakaoC := make(chan []kakaoExtractedJob)
	go jobscrapper.KakaoCrawling(kakaoC)
	kakaoJobs := <-kakaoC

	contents := etc.StructToStr(&kakaoJobs)
	fmt.Println(contents)

	subscribers := aws.GetSubscribers()
	fmt.Println(subscribers)
	aws.SendMail(contents, subscribers)
}
