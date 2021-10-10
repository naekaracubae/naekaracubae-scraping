package main

import (
	"fmt"
	jobscrapper2 "github.com/msyhu/GobbyIsntFree/developerilbo/jobscrapper"
	"testing"
)

func TestJobscrapping(t *testing.T) {

	kakaoC := make(chan []kakaoExtractedJob)
	lineC := make(chan []lineExtractedJob)
	go jobscrapper2.LineCrawling(lineC)
	go jobscrapper2.KakaoCrawling(kakaoC)
	kakaoJobs := <-kakaoC
	lineJobs := <-lineC

	fmt.Println(kakaoJobs)
	fmt.Println(lineJobs)
}
