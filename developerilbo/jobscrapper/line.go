package jobscrapper

import (
	_ "fmt"
	"github.com/PuerkitoBio/goquery"
	etc2 "github.com/msyhu/GobbyIsntFree/developerilbo/etc"
	_struct2 "github.com/msyhu/GobbyIsntFree/developerilbo/struct"
	"net/http"
	_ "strconv"
	_ "strings"
)

var linebaseURL string = "https://careers.linecorp.com/ko/jobs?co=Korea&ca=Engineering"

type lineJob = _struct2.Line

func LineCrawling(lineC chan<- []lineJob) {
	var jobs []lineJob
	c := make(chan []lineJob)

	go LineGetPage(c)

	lineJobs := <-c
	jobs = append(jobs, lineJobs...)

	lineC <- jobs
}

func LineGetPage(mainC chan<- []lineJob) {
	var jobs []lineJob
	c := make(chan lineJob)
	res, err := http.Get(linebaseURL)
	etc2.CheckErr(err)
	etc2.CheckCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	etc2.CheckErr(err)

	searchCards := doc.Find(".list_jobs>li")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go LineExtractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs

}

func LineExtractJob(card *goquery.Selection, c chan<- lineJob) {
	// title
	title := card.Find(".tit_jobs").Text()

	// Url, Id
	link, _ := card.Find(".link_jobs").Attr("href")
	fullLink := "https://careers.line.com" + link
	id := extractId(link)

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

	c <- lineJob{Title: title, EndDate: endDate, Location: location, Company: company, JobType: jobType, Url: fullLink, Id: id}

}
