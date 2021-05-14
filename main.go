package main

import (
	"encoding/json"
	"fmt"
	"github.com/msyhu/GobbyIsntFree/etc"
	"github.com/msyhu/GobbyIsntFree/kakaoCrawler"
	"strconv"
	"strings"
)

type kakaoExtractedJob = kakaoCrawler.ExtractedJob

func main() {

	startCrawling()
}

func startCrawling() {
	kakaoC := make(chan []kakaoExtractedJob)
	go kakaoCrawler.Crawling(kakaoC)
	kakaoJobs := <-kakaoC
	fmt.Println(kakaoJobs)

	// 모듈화
	var contents strings.Builder
	for idx, kakaoJob := range kakaoJobs {
		jsonBytes, err := json.Marshal(kakaoJob)
		etc.CheckErr(err)
		jsonString := string(jsonBytes)
		idxString := strconv.Itoa(idx) + ". " + jsonString
		contents.WriteString(idxString)
		contents.WriteString("</br>")
	}

	subscribers := etc.GetSubscribers()
	etc.SendMail(contents.String(), subscribers)
}
