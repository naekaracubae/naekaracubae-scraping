package jobscrapper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	etc2 "github.com/msyhu/GobbyIsntFree/developerilbo/etc"
	_struct "github.com/msyhu/GobbyIsntFree/developerilbo/struct"
	"net/http"
	"strings"
)

var linebaseURL string = "https://careers.linecorp.com/ko/jobs"

type lineJob = _struct.Line

func LineCrawling(lineC chan<- []lineJob) {
	//var jobs []lineJob

	res, err := http.Get(linebaseURL)
	etc2.CheckErr(err)
	etc2.CheckCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	etc2.CheckErr(err)

	searchCards := doc.Find(".job_list>li")

	c := make(chan lineJob)
	searchCards.Each(func(i int, card *goquery.Selection) {
		go LineExtractJob(card, c)
	})
}

func LineExtractJob(card *goquery.Selection, c chan<- lineJob) {
	// infos
	infos := card.Find(".text_filter").Text()
	if strings.Contains(infos, "Engineering") == false {
		return
	}
	fmt.Println(infos)

	// infos -> LOCATION, COMPANY 분리

	//fmt.Println(infos)

	// title
	//title := card.Find(".title").Text()
	//fmt.Println(title)

	//// Url, Id
	//link, _ := card.Find(".link_jobs").Attr("href")
	//fullLink := "https://careers.line.com" + link
	//id := extractId(link)
	//
	//// endDate, location
	//var endDateAndLocation []string
	//card.Find(".list_info>dd").Each(func(i int, s *goquery.Selection) {
	//	endDateAndLocation = append(endDateAndLocation, s.Text())
	//})
	//
	//var endDate = ""
	//var location = ""
	//if len(endDateAndLocation) == 2 {
	//	endDate = endDateAndLocation[0]
	//	location = endDateAndLocation[1]
	//} else {
	//	endDate = endDateAndLocation[0]
	//}
	//
	////jobGroups
	//var jobGroups []string
	//card.Find(".list_tag>a").Each(func(i int, s *goquery.Selection) {
	//	jobGroup, _ := s.Attr("data-code")
	//	jobGroups = append(jobGroups, jobGroup)
	//})
	////company, jobType
	//var companyAndJobType []string
	//card.Find(".item_subinfo>dd").Each(func(i int, s *goquery.Selection) {
	//	companyAndJobType = append(companyAndJobType, s.Text())
	//})
	//company := companyAndJobType[0]
	//jobType := companyAndJobType[1]
	//
	//c <- kakaoJob{Title: title, EndDate: endDate, Location: location, JobGroups: jobGroups, Company: company, JobType: jobType, Url: fullLink, Id: id}

}

//
//func extractId(link string) string {
//	l1 := strings.Split(link, "?")[0]
//	l2 := strings.Split(l1, "/")[2]
//
//	return l2
//}
