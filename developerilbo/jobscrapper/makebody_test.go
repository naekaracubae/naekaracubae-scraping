package jobscrapper_test

import (
	"fmt"
	jobscrapper2 "github.com/msyhu/GobbyIsntFree/developerilbo/jobscrapper"
	"testing"
)

func TestMakeHtmlBody(t *testing.T) {
	contents := jobscrapper2.MakeHtmlBody()

	fmt.Println(contents)
}
