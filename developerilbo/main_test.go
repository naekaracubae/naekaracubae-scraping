package main

import (
	"fmt"
	jobscrapper2 "github.com/msyhu/GobbyIsntFree/developerilbo/jobscrapper"
	_struct "github.com/msyhu/GobbyIsntFree/developerilbo/struct"
	"testing"
)

type lineJob = _struct.Line

func TestJobscrapping(t *testing.T) {

	kakaoC := make(chan []kakaoExtractedJob)
	go jobscrapper2.KakaoCrawling(kakaoC)
	kakaoJobs := <-kakaoC

	var lineJobs []lineJob
	lineJobs = jobscrapper2.LineCrawling()

	fmt.Println(kakaoJobs)
	fmt.Println(lineJobs)
}
