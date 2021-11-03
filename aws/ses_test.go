package aws_test

import (
	aws2 "github.com/msyhu/naekaracubae-scraping/aws"
	"github.com/msyhu/naekaracubae-scraping/jobscrapper"
	_struct2 "github.com/msyhu/naekaracubae-scraping/struct"
	"testing"
)

func Test_Sendmail(t *testing.T) {
	contents := jobscrapper.MakeHtmlBody()
	subscriber := _struct2.Subscriber{"msyhu", "msyhu@korea.ac.kr"}
	subscriber2 := _struct2.Subscriber{"msyhu2", "anstkd07@gmail.com"}
	var subscribers []_struct2.Subscriber
	subscribers = append(subscribers, subscriber)
	subscribers = append(subscribers, subscriber2)

	aws2.SendMail(contents, subscribers)
}
