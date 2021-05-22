package aws_test

import (
	"github.com/msyhu/GobbyIsntFree/aws"
	"testing"
)

func TestSendmail(t *testing.T) {
	contents := "hello world"
	aws.SendMail(&contents, nil)
}
