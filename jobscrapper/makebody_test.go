package jobscrapper_test

import (
	"fmt"
	"github.com/msyhu/GobbyIsntFree/jobscrapper"
	"testing"
)

func TestMakeHtmlBody(t *testing.T) {
	contents := jobscrapper.MakeHtmlBody()

	fmt.Println(contents)
}
