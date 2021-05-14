package etc_test

import (
	"github.com/msyhu/GobbyIsntFree/etc"
	"testing"
)

func TestGetSubscribers(t *testing.T) {
	etc.GetSubscribers()
}

func TestGetSenders(t *testing.T) {
	etc.GetSenders()
}
