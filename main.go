package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/msyhu/GobbyIsntFree/etc"
	"github.com/msyhu/GobbyIsntFree/kakaoCrawler"
	"strconv"
	"strings"
)

type kakaoExtractedJob = kakaoCrawler.ExtractedJob

func main() {
	lambda.Start(startCrawling)
}

func startCrawling() {
	kakaoC := make(chan []kakaoExtractedJob)
	go kakaoCrawler.Crawling(kakaoC)
	kakaoJobs := <-kakaoC
	fmt.Println(kakaoJobs)

	var contents strings.Builder
	for idx, kakaoJob := range kakaoJobs {
		jsonBytes, err := json.Marshal(kakaoJob)
		etc.CheckErr(err)
		jsonString := string(jsonBytes)
		idxString := strconv.Itoa(idx) + ". " + jsonString
		contents.WriteString(idxString)
		contents.WriteString("</br>")
	}

	etc.SendMail(contents.String())
}
