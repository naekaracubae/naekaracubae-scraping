package jobscrapper_test

import (
	"fmt"
	jobscrapper2 "github.com/msyhu/GobbyIsntFree/developerilbo/jobscrapper"
	_struct2 "github.com/msyhu/GobbyIsntFree/developerilbo/struct"
	"testing"
)

type kakaoJob = _struct2.Kakao

func TestKakaoGetPages(t *testing.T) {
	pages := jobscrapper2.KakaoGetPages()

	if pages != 14 {
		t.Error("Wrong result", pages)
	}
}

func TestKakaoGetPage(t *testing.T) {
	c := make(chan []kakaoJob)
	jobscrapper2.KakaoGetPage(1, c)

}

func TestKakaoCrawling(t *testing.T) {
	var jobs []kakaoJob
	c := make(chan []kakaoJob)

	totalPages := jobscrapper2.KakaoGetPages()
	for i := 1; i <= totalPages; i++ {
		go jobscrapper2.KakaoGetPage(i, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	for _, job := range jobs {
		fmt.Println(job)
	}

}
