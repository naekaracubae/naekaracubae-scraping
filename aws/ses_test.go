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
	var subscribers []_struct2.Subscriber
	subscribers = append(subscribers, subscriber)

	aws2.SendMail(contents, subscribers)
}
