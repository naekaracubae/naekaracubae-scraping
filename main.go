package main

import (
	"encoding/json"
	"fmt"
	"github.com/msyhu/GobbyIsntFree/etc"
	"github.com/msyhu/GobbyIsntFree/kakaoCrawler"
	"github.com/robfig/cron/v3"
	"os"
	"os/signal"
	"strconv"
	"strings"
)

type kakaoExtractedJob = kakaoCrawler.ExtractedJob

func main() {
	c := cron.New()
	c.AddFunc("@midnight", startCrawling)
	go c.Start()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
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
