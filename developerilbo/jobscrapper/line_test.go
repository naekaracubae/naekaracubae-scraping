package jobscrapper

import (
	//"fmt"
	"testing"
)

func TestLineCrawling(t *testing.T) {
	//var jobs []lineJob
	c := make(chan []lineJob)
	LineCrawling(c)

}
