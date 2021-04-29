package main

import (
	"fmt"
	"github.com/msyhu/GobbyIsntFree/indeedCrawler"
)

type extractedJob = indeedCrawler.ExtractedJob

func main() {
	indeedC := make(chan []extractedJob)
	go indeedCrawler.Crawling(indeedC)
	indeedJobs := <-indeedC
	for _, indeedJob := range indeedJobs {
		fmt.Println(indeedJob)
	}

}
