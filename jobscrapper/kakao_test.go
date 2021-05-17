package jobscrapper_test

import (
	"fmt"
	"github.com/msyhu/GobbyIsntFree/jobscrapper"
	_struct "github.com/msyhu/GobbyIsntFree/struct"
	"testing"
)

type kakaoJob = _struct.Kakao

func TestKakaoGetPages(t *testing.T) {
	pages := jobscrapper.KakaoGetPages()

	if pages != 14 {
		t.Error("Wrong result", pages)
	}
}

func TestKakaoGetPage(t *testing.T) {
	c := make(chan []kakaoJob)
	jobscrapper.KakaoGetPage(1, c)

}

func TestKakaoCrawling(t *testing.T) {
	var jobs []kakaoJob
	c := make(chan []kakaoJob)

	totalPages := jobscrapper.KakaoGetPages()
	for i := 1; i <= totalPages; i++ {
		go jobscrapper.KakaoGetPage(i, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	for _, job := range jobs {
		fmt.Println(job)
	}

}
