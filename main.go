package main

import (
	"fmt"
	"github.com/msyhu/GobbyIsntFree/crawlers"
)

type extractedJob = crawlers.IndeedExtractedJob

func main() {
	indeedC := make(chan []extractedJob)
	go crawlers.CrawlIndeed(indeedC)
	indeedJobs := <-indeedC
	for _, indeedJob := range indeedJobs {
		fmt.Println(indeedJob)
	}

}
