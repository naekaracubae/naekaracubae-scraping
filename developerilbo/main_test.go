package main

import (
	"github.com/msyhu/naekaracubae-scraping/developerilbo/jobscrapper"
	"log"
	"testing"
)

func Test_전체_회사_스크래핑(t *testing.T) {
	log.Println("전체_회사_스크래핑 start")

	kakaoC := make(chan []kakaoJob)
	lineC := make(chan []lineJob)
	go jobscrapper.LineCrawling(lineC)
	go jobscrapper.KakaoCrawling(kakaoC)
	kakaoJobs := <-kakaoC
	lineJobs := <-lineC

	log.Println(kakaoJobs)
	log.Println(lineJobs)
	log.Println("전체_회사_스크래핑 finish")
}
