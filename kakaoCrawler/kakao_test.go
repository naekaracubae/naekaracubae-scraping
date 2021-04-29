package kakaoCrawler_test

import (
	"github.com/msyhu/GobbyIsntFree/kakaoCrawler"
	"testing"
)

func TestGetPages(t *testing.T) {
	pages := kakaoCrawler.GetPages()

	if pages != 7 {
		t.Error("Wrong result", pages)
	}
}
