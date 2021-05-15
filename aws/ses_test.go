package aws_test

import (
	"github.com/msyhu/GobbyIsntFree/aws"
	"testing"
)

func TestSendmail(t *testing.T) {
	aws.SendMail("hello world", nil)
}
