package jobscrapper_test

import (
	"fmt"
	"github.com/msyhu/GobbyIsntFree/jobscrapper"
	"testing"
)

type extractedJob = jobscrapper.ExtractedJob

func TestGetPages(t *testing.T) {
	pages := jobscrapper.GetPages()

	if pages != 14 {
		t.Error("Wrong result", pages)
	}
}

func TestGetPage(t *testing.T) {
	c := make(chan []extractedJob)
	jobscrapper.GetPage(1, c)

}

func TestCrawling(t *testing.T) {
	var jobs []extractedJob
	c := make(chan []extractedJob)

	totalPages := jobscrapper.GetPages()
	for i := 1; i <= totalPages; i++ {
		go jobscrapper.GetPage(i, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	for _, job := range jobs {
		fmt.Println(job)
	}

}
