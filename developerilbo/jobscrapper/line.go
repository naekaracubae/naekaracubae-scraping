package jobscrapper

import (
	"github.com/PuerkitoBio/goquery"
	etc "github.com/msyhu/naekaracubae-scraping/developerilbo/etc"
	_struct "github.com/msyhu/naekaracubae-scraping/developerilbo/struct"
	"log"
	"net/http"
	"strings"
)

var linebaseURL = "https://careers.linecorp.com/ko/jobs"

type lineJob = _struct.Line

// 한 페이지 단위 전체 직무 스크래핑
func LineCrawling(lineC chan<- []lineJob) {
	log.Println("LineCrawling start")
	var jobs []lineJob

	res, err := http.Get(linebaseURL)
	etc.CheckErr(err)
	etc.CheckCode(res)
	log.Println("response : ", res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	etc.CheckErr(err)

	searchCards := doc.Find(".job_list>li")

	c := make(chan lineJob)
	searchCards.Each(func(i int, card *goquery.Selection) {
		go LineExtractJob(card, c, i)
	})

	// Engineering만 뽑아내서 그 횟수만큼만 반복문을 통해 채널에서 빼내야 한다.
	// 이렇게 필터링 안 하면 계속 채널에서 뽑아내려고 하기 때문에 프로그램이 종료 안 됨.
	countEngineering := 0
	searchCards.Each(func(i int, card *goquery.Selection) {
		infos := card.Find(".text_filter").Text()
		if strings.Contains(infos, "Engineering") {
			countEngineering += 1
		}
	})

	for i := 0; i < countEngineering; i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	lineC <- jobs
}

// 직무 하나 단위 스크래핑
func LineExtractJob(card *goquery.Selection, c chan<- lineJob, idx int) {
	log.Println(idx, "th LineExtractJob start")
	// infos
	infos := card.Find(".text_filter").Text()
	if strings.Contains(infos, "Engineering") == false {
		return
	}
	//fmt.Println(infos)

	// infos -> LOCATION, COMPANY 분리
	splitByJob := strings.Split(infos, "Engineering")
	splitByBar := strings.Split(splitByJob[0], "|")
	location := strings.Trim(splitByBar[0], " ")
	company := strings.Trim(splitByBar[1], " ")
	//fmt.Println(location, ",", company)

	// title
	title := strings.Trim(card.Find(".title").Text(), " ")
	//fmt.Println(location, ",", company, ",", title)

	// startdate, enddate
	date := card.Find(".date").Text()
	//fmt.Println(location, ",", company, ",", title, ",", date)
	splitByStartEnd := strings.Split(date, "~")
	startDate := strings.Trim(splitByStartEnd[0], " ")
	endDate := strings.Trim(splitByStartEnd[1], " ")
	//fmt.Println(location, ",", company, ",", title, ",", startDate, ",", endDate)

	// url
	link, _ := card.Find("a").Attr("href")
	fullLink := "https://careers.linecorp.com" + link
	//fmt.Println(location, ",", company, ",", title, ",", startDate, ",", endDate, ",", fullLink)
	//fmt.Println(fullLink)

	// id
	id := strings.Split(link, "/")[3]
	log.Println(idx, "th result : [",
		location, ",",
		company, ",",
		title, ",",
		startDate, ",",
		endDate, ",",
		fullLink, ",",
		id, "]")

	c <- lineJob{
		Title:     title,
		StartDate: startDate,
		EndDate:   endDate,
		Location:  location,
		Company:   company,
		Url:       fullLink,
		Id:        id}

}
