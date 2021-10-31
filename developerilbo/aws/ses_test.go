package aws_test

import (
	aws2 "github.com/msyhu/naekaracubae-scraping/developerilbo/aws"
	"github.com/msyhu/naekaracubae-scraping/developerilbo/jobscrapper"
	_struct2 "github.com/msyhu/naekaracubae-scraping/developerilbo/struct"
	"testing"
)

func TestSendmail(t *testing.T) {
	contents := jobscrapper.MakeHtmlBody()
	subscriber := _struct2.Subscriber{"msyhu", "msyhu@korea.ac.kr"}
	var subscribers []_struct2.Subscriber
	subscribers = append(subscribers, subscriber)

	aws2.SendMail(contents, subscribers)
}
