package aws_test

import (
	aws2 "github.com/msyhu/GobbyIsntFree/developerilbo/aws"
	"testing"
)

func TestSendmail(t *testing.T) {
	contents := "hello world"
	aws2.SendMail(&contents, nil)
}
