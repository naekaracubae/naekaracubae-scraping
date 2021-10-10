package jobscrapper

import (
	//"fmt"
	"fmt"
	"testing"
)

func TestLineCrawling(t *testing.T) {
	var jobs []lineJob
	jobs = LineCrawling()

	fmt.Println(jobs)

}

//func TestLineExtractJob(t *testing.T) {
//	var jobs []lineJob
//	c := make(chan []lineJob)
//
//	go LineExtractJob(c)
//
//}
