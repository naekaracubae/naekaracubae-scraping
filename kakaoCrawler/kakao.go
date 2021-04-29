package kakaoCrawler

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/msyhu/GobbyIsntFree/etc"
	"net/http"
)

var baseURL string = "https://careers.kakao.com/jobs?part=TECHNOLOGY"

func GetPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	etc.CheckErr(err)
	etc.CheckCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	etc.CheckErr(err)

	doc.Find(".paging_list").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	// 양쪽 화살표 4개 빼주고 현재 페이지 1 더해줌
	return pages - 4 + 1
}
