package etc_test

import (
	"github.com/msyhu/GobbyIsntFree/etc"
	"testing"
)

func TestSendmail(t *testing.T) {
	etc.SendMail("hello world", nil)
}
