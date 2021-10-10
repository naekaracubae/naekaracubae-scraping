package jobscrapper

import (
	"github.com/PuerkitoBio/goquery"
	etc2 "github.com/msyhu/GobbyIsntFree/developerilbo/etc"
	_struct "github.com/msyhu/GobbyIsntFree/developerilbo/struct"
	"net/http"
	"strings"
)

var linebaseURL = "https://careers.linecorp.com/ko/jobs"

type lineJob = _struct.Line

func LineCrawling(lineC chan<- []lineJob) {
	//fmt.Println("line start!!")
	var jobs []lineJob

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

	countEngineering := 0
	searchCards.Each(func(i int, card *goquery.Selection) {
		infos := card.Find(".text_filter").Text()
		if strings.Contains(infos, "Engineering") {
			countEngineering += 1
		}
	})

	for i := 0; i < countEngineering; i++ {
		job := <-c
		//fmt.Println(job)
		jobs = append(jobs, job)
	}

	lineC <- jobs
}

func LineExtractJob(card *goquery.Selection, c chan<- lineJob) {
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
	//fmt.Println(location, ",", company, ",", title, ",", startDate, ",", endDate, ",", fullLink, ",", id)

	c <- lineJob{
		Title:     title,
		StartDate: startDate,
		EndDate:   endDate,
		Location:  location,
		Company:   company,
		Url:       fullLink,
		Id:        id}

}
