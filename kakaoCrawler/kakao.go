package kakaoCrawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/msyhu/GobbyIsntFree/etc"
	"net/http"
	"strconv"
)

var baseURL string = "https://careers.kakao.com/jobs?part=TECHNOLOGY&company=ALL"

// TODO : 페이지 끝까지 가서 가져와야 함.
// 페이지 수를 가져온다
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

// 하나의 페이지에서 직무를 가져와서 하나씩 채널로 넘겨준다.
func GetPage(page int) {

	pageURL := baseURL + "&page=" + strconv.Itoa(page)
	fmt.Println(pageURL)
	res, err := http.Get(pageURL)
	etc.CheckErr(err)
	etc.CheckCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	etc.CheckErr(err)

	searchCards := doc.Find(".list_jobs>li")

	searchCards.Each(func(i int, card *goquery.Selection) {
		extractJob(card)
	})
	//
	//for i := 0; i < searchCards.Length(); i++ {
	//	job := <-c
	//	jobs = append(jobs, job)
	//}
	//
	//mainC <- jobs

}

func extractJob(card *goquery.Selection) {
	// title
	title := card.Find(".tit_jobs").Text()

	// endDate, location
	var endDateAndLocation []string
	card.Find(".list_info>dd").Each(func(i int, s *goquery.Selection) {
		// I don't want the first string into my array, so I filter it
		endDateAndLocation = append(endDateAndLocation, s.Text())
	})
	//jobGroups
	var jobGroups []string
	card.Find(".list_tag>a").Each(func(i int, s *goquery.Selection) {
		// I don't want the first string into my array, so I filter it
		jobGroup, _ := s.Attr("data-code")
		jobGroups = append(jobGroups, jobGroup)
	})
	//company, jobType
	var companyAndJobType []string
	card.Find(".item_subinfo>dd").Each(func(i int, s *goquery.Selection) {
		// I don't want the first string into my array, so I filter it
		companyAndJobType = append(companyAndJobType, s.Text())
	})

	//c <- extractedJob{id, title, company, location}

	fmt.Println(title, endDateAndLocation, jobGroups, companyAndJobType)

}
