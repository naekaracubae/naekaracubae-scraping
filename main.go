package main

import (
	"fmt"
	"github.com/msyhu/GobbyIsntFree/indeedCrawlers"
)

type extractedJob = indeedCrawlers.ExtractedJob

func main() {
	indeedC := make(chan []extractedJob)
	go indeedCrawlers.Crawling(indeedC)
	indeedJobs := <-indeedC
	for _, indeedJob := range indeedJobs {
		fmt.Println(indeedJob)
	}

}
