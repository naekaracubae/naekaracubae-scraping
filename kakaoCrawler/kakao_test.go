package kakaoCrawler_test

import (
	"fmt"
	"github.com/msyhu/GobbyIsntFree/kakaoCrawler"
	"testing"
)

type extractedJob = kakaoCrawler.ExtractedJob

func TestGetPages(t *testing.T) {
	pages := kakaoCrawler.GetPages()

	if pages != 7 {
		t.Error("Wrong result", pages)
	}
}

func TestGetPage(t *testing.T) {
	c := make(chan []extractedJob)
	kakaoCrawler.GetPage(1, c)

}

func TestCrawling(t *testing.T) {
	var jobs []extractedJob
	c := make(chan []extractedJob)

	totalPages := kakaoCrawler.GetPages()
	for i := 0; i < totalPages; i++ {
		go kakaoCrawler.GetPage(i, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	for _, job := range jobs {
		fmt.Println(job)
	}

}
