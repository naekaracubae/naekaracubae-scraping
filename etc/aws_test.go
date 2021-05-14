package etc_test

import (
	"fmt"
	"github.com/msyhu/GobbyIsntFree/etc"
	"testing"
)

func TestGetSubscribers(t *testing.T) {
	etc.GetSubscribers()
}

func TestGetSenders(t *testing.T) {
	sender := etc.GetSenders()
	fmt.Println(sender)
}
