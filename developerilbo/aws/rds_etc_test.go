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

func Test_GetSecretForKakao(t *testing.T) {
	gobbyRdsSecret := aws2.GetSecret()
	log.Println(gobbyRdsSecret)

	if gobbyRdsSecret == nil {
		t.Error("Wrong result")
	}
}
