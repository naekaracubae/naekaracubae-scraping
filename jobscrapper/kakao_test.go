package jobscrapper_test

import (
	"fmt"
	"github.com/msyhu/naekaracubae-scraping/jobscrapper"
	"github.com/msyhu/naekaracubae-scraping/struct"
	"log"
	"testing"
)

type kakaoJob = _struct.Kakao

func Test_카카오_전체_페이지수(t *testing.T) {
	log.Println("Test_카카오_전체_페이지수 start")
	pages := jobscrapper.KakaoGetPages()
	fmt.Println(pages)

	if pages != 16 {
		t.Error("Wrong result", pages)
	}
	log.Println("Test_카카오_전체_페이지수 finish")
}

func Test_카카오_한페이지_스크래핑(t *testing.T) {
	log.Println("Test_카카오_한페이지_스크래핑 start")
	c := make(chan []kakaoJob)
	jobscrapper.KakaoGetPage(1, c)
	log.Println("Test_카카오_한페이지_스크래핑 finish")
}

func Test_카카오_전체_스크래핑(t *testing.T) {
	log.Println("Test_카카오_전체_스크래핑 start")
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
	log.Println("Test_카카오_전체_스크래핑 finish")
}
