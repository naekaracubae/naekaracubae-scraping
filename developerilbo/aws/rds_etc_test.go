package aws_test

import (
	aws2 "github.com/msyhu/naekaracubae-scraping/developerilbo/aws"
	"log"
	"testing"
)

func Test_GetSubscribers(t *testing.T) {
	subscribers := aws2.GetSubscribers()
	log.Println(subscribers)
}
