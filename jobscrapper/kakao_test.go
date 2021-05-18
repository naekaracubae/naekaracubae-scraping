package jobscrapper_test

import (
	"fmt"
	"github.com/msyhu/GobbyIsntFree/aws"
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

func TestCheckJob(t *testing.T) {
	testKakaoStruct := kakaoJob{
		Title:    "test",
		EndDate:  "채용시까지",
		Location: "판교",
		Company:  "kakao",
		Url:      "https://careers.kakao.com/jobs/P-9349?part=TECHNOLOGY&company=ALL",
		JobType:  "정규직",
		Id:       "P-9349",
	}

	aws.CheckJob(&testKakaoStruct)
}
