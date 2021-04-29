package indeedCrawler

import (
	"encoding/csv"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/msyhu/GobbyIsntFree/etc"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

type ExtractedJob struct {
	id       string
	title    string
	company  string
	location string
}

type extractedJob = ExtractedJob

func Crawling(indeedC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan []extractedJob)

	totalPages := getPages()
	for i := 0; i < totalPages; i++ {
		go getPage(i, c)
	}

	for i := 0; i < totalPages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	indeedC <- jobs
}

func getPage(page int, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println(pageURL)
	res, err := http.Get(pageURL)
	etc.CheckErr(err)
	etc.CheckCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	etc.CheckErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")

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
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".title>a").Text())
	company := cleanString(card.Find(".sjcl>div>span").Text())
	location := cleanString(card.Find(".sjcl").Text())
	c <- extractedJob{id, title, company, location}

}

// csv로 저장
func writeJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	etc.CheckErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	header := []string{"id", "title", "company", "location"}
	wErr := w.Write(header)
	etc.CheckErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{job.id, job.title, job.company, job.location}
		jwErr := w.Write(jobSlice)
		etc.CheckErr(jwErr)
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	etc.CheckErr(err)
	etc.CheckCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	etc.CheckErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}
