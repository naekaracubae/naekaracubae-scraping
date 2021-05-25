package jobscrapper

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	etc2 "github.com/msyhu/GobbyIsntFree/developerilbo/etc"
	_struct2 "github.com/msyhu/GobbyIsntFree/developerilbo/struct"
	"net/http"
	"strconv"
)

var indeedbaseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

type indeedJob = _struct2.Indeed

func IndeedCrawling(indeedC chan<- []indeedJob) {
	var jobs []indeedJob
	c := make(chan []indeedJob)

	totalPages := IndeedGetPages()
	for i := 0; i < totalPages; i++ {
		go IndeedGetPage(i, c)
	}

	for i := 0; i < totalPages; i++ {
		indeedJobs := <-c
		fmt.Println(indeedJobs)
		jobs = append(jobs, indeedJobs...)
	}

	//indeedC <- jobs
}

func IndeedGetPage(page int, mainC chan<- []indeedJob) {
	var jobs []indeedJob
	c := make(chan indeedJob)
	pageURL := indeedbaseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println(pageURL)
	res, err := http.Get(pageURL)
	etc2.CheckErr(err)
	etc2.CheckCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	etc2.CheckErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")

	searchCards.Each(func(i int, card *goquery.Selection) {
		go IndeedExtractJob(card, c)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs

}

func IndeedExtractJob(card *goquery.Selection, c chan<- indeedJob) {
	id, _ := card.Attr("data-jk")
	title := etc2.CleanString(card.Find(".title>a").Text())
	company := etc2.CleanString(card.Find(".sjcl>div>span").Text())
	location := etc2.CleanString(card.Find(".sjcl").Text())

	//fmt.Println(company)

	c <- indeedJob{id, title, company, location}

}

func IndeedGetPages() int {
	pages := 0
	res, err := http.Get(indeedbaseURL)
	etc2.CheckErr(err)
	etc2.CheckCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	etc2.CheckErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}
