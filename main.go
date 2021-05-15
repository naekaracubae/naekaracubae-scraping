package main

import (
	"github.com/msyhu/GobbyIsntFree/aws"
	"github.com/msyhu/GobbyIsntFree/etc"
	"github.com/msyhu/GobbyIsntFree/kakaoCrawler"
	_struct "github.com/msyhu/GobbyIsntFree/struct"
)

type kakaoExtractedJob = _struct.Kakao

func main() {
	start()
}

func start() {
	kakaoC := make(chan []kakaoExtractedJob)
	go kakaoCrawler.Crawling(kakaoC)
	kakaoJobs := <-kakaoC

	contents := etc.StructToStr(&kakaoJobs)
	subscribers := aws.GetSubscribers()
	aws.SendMail(contents, subscribers)
}
