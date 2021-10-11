package jobscrapper

import (
	"github.com/PuerkitoBio/goquery"
	etc "github.com/msyhu/GobbyIsntFree/developerilbo/etc"
	_struct "github.com/msyhu/GobbyIsntFree/developerilbo/struct"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var kakaobaseURL = "https://careers.kakao.com/jobs?part=TECHNOLOGY&company=ALL"

type kakaoJob = _struct.Kakao

// 카카오 크롤링 수행하는 메인문
func KakaoCrawling(kakaoC chan<- []kakaoJob) {
	log.Println("KakaoCrawling start")

	var jobs []kakaoJob
	c := make(chan []kakaoJob)

	totalPages := KakaoGetPages()

	for i := 1; i <= totalPages; i++ {
		go KakaoGetPage(i, c)
	}

	for i := 0; i < totalPages; i++ {
		kakaoJobs := <-c
		jobs = append(jobs, kakaoJobs...)
	}

	kakaoC <- jobs
	log.Println("KakaoCrawling finished")
}

// 전체 페이지 수를 가져온다
func KakaoGetPages() int {
	res, err := http.Get(kakaobaseURL)
	etc.CheckErr(err)
	etc.CheckCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	etc.CheckErr(err)

	pageSelection := doc.Find(".paging_list").Find("a")
	lastPageHref, _ := pageSelection.Last().Attr("href")
	lastPage := strings.Split(lastPageHref, "=")[1]
	page, err := strconv.Atoi(lastPage)
	etc.CheckErr(err)

	return page
}

// 한 페이지 단위 전체 직무 스크래핑
func KakaoGetPage(page int, mainC chan<- []kakaoJob) {
	log.Println(page, "page KakaoGetPage start")

	var jobs []kakaoJob
	c := make(chan kakaoJob)
	pageURL := kakaobaseURL + "&page=" + strconv.Itoa(page)
	res, err := http.Get(pageURL)
	etc.CheckErr(err)
	etc.CheckCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	etc.CheckErr(err)

	searchCards := doc.Find(".list_jobs>li")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go KakaoExtractJob(card, c, i, page)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs

}

// 직무 하나 단위 스크래핑
func KakaoExtractJob(card *goquery.Selection, c chan<- kakaoJob, idx int, page int) {
	log.Println(page, "page ", idx, "th KakaoExtractJob start")

	// title
	title := card.Find(".tit_jobs").Text()

	// Url, Id
	link, _ := card.Find(".link_jobs").Attr("href")
	fullLink := "https://careers.kakao.com" + link
	id := extractId(link)
	//fmt.Println(link)

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

	log.Println(page, "page ", idx, "th result : [",
		title, ",",
		endDate, ",",
		location, ",",
		jobGroups, ",",
		company, ",",
		fullLink, ",",
		id, "]")

	c <- kakaoJob{
		Title:     title,
		EndDate:   endDate,
		Location:  location,
		JobGroups: jobGroups,
		Company:   company,
		JobType:   jobType,
		Url:       fullLink,
		Id:        id}

}

func extractId(link string) string {
	l1 := strings.Split(link, "?")[0]
	l2 := strings.Split(l1, "/")[2]

	return l2
}
