package kakaoCrawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/msyhu/GobbyIsntFree/etc"
	"net/http"
	"strconv"
)

var baseURL string = "https://careers.kakao.com/jobs?part=TECHNOLOGY&company=ALL"

type ExtractedJob struct {
	title     string
	endDate   string
	location  string
	jobGroups []string
	company   string
	jobType   string
}

type extractedJob = ExtractedJob

func Crawling(kakaoC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan []extractedJob)

	totalPages := GetPages()
	for i := 0; i < totalPages; i++ {
		go GetPage(i, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	kakaoC <- jobs
}

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
func GetPage(page int, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
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
		go extractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs

}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	// title
	title := card.Find(".tit_jobs").Text()

	// endDate, location
	var endDateAndLocation []string
	card.Find(".list_info>dd").Each(func(i int, s *goquery.Selection) {
		endDateAndLocation = append(endDateAndLocation, s.Text())
	})

	var endDate = ""
	var location = ""
	if len(endDateAndLocation) == 2 {
		endDate = endDateAndLocation[0]
		location = endDateAndLocation[1]
	} else {
		endDate = endDateAndLocation[0]
	}

	//jobGroups
	var jobGroups []string
	card.Find(".list_tag>a").Each(func(i int, s *goquery.Selection) {
		jobGroup, _ := s.Attr("data-code")
		jobGroups = append(jobGroups, jobGroup)
	})
	//company, jobType
	var companyAndJobType []string
	card.Find(".item_subinfo>dd").Each(func(i int, s *goquery.Selection) {
		companyAndJobType = append(companyAndJobType, s.Text())
	})
	company := companyAndJobType[0]
	jobType := companyAndJobType[1]

	c <- extractedJob{title, endDate, location, jobGroups, company, jobType}

}
