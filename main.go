package main

import (
	"fmt"
	"github.com/msyhu/GobbyIsntFree/indeedCrawler"
	"github.com/msyhu/GobbyIsntFree/kakaoCrawler"
)

type indeedExtractedJob = indeedCrawler.ExtractedJob
type kakaoExtractedJob = kakaoCrawler.ExtractedJob

func main() {
	indeedC := make(chan []indeedExtractedJob)
	kakaoC := make(chan []kakaoExtractedJob)
	go indeedCrawler.Crawling(indeedC)
	go kakaoCrawler.Crawling(kakaoC)
	indeedJobs := <-indeedC
	kakaoJobs := <-kakaoC

	for _, indeedJob := range indeedJobs {
		fmt.Println(indeedJob)
	}

	for _, kakaoJob := range kakaoJobs {
		fmt.Println(kakaoJob)
	}

}
