package main

import (
	jobscrapper2 "github.com/msyhu/naekaracubae-scraping/jobscrapper"
	"log"
	"testing"
)

func Test_전체_회사_스크래핑(t *testing.T) {
	log.Println("전체_회사_스크래핑 start")

	kakaoC := make(chan []kakaoJob)
	lineC := make(chan []lineJob)
	go jobscrapper2.LineCrawling(lineC)
	go jobscrapper2.KakaoCrawling(kakaoC)
	kakaoJobs := <-kakaoC
	lineJobs := <-lineC

	log.Println(kakaoJobs)
	log.Println(lineJobs)
	log.Println("전체_회사_스크래핑 finish")
}
