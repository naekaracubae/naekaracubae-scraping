package jobscrapper

import (
	"log"
	"testing"
)

// 라인은 한페이지가 전부이므로 이것만 테스트
func Test_라인_전체_스크래핑(t *testing.T) {
	log.Println("Test_라인_전체_스크래핑 start")

	lineC := make(chan []lineJob)
	go LineCrawling(lineC)
	lineJobs := <-lineC

	log.Println(lineJobs)
	log.Println("Test_라인_전체_스크래핑 finish")
}
