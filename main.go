package main

import (
	"encoding/json"
	"fmt"
	"github.com/msyhu/GobbyIsntFree/etc"
	"github.com/msyhu/GobbyIsntFree/kakaoCrawler"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"strconv"
	"strings"
)

type kakaoExtractedJob = kakaoCrawler.ExtractedJob

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func main() {
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
	//
	//e, err := json.Marshal(&kakaoJobs)
	//etc.CheckErr(err)
	//contents += string(e)

	//out, err := json.Marshal(kakaoJobs)
	//etc.CheckErr(err)
	// 메일 보내기
	//etc.SendMail(out)

	userJson, err := ioutil.ReadFile("./secrets/sendmail.json") // articles.json 파일의 내용을 읽어서 바이트 슬라이스에 저장
	etc.CheckErr(err)
	var user User
	json.Unmarshal(userJson, &user)

	subscribersJson, err := ioutil.ReadFile("./secrets/subscribers.json") // articles.json 파일의 내용을 읽어서 바이트 슬라이스에 저장
	etc.CheckErr(err)
	var subscribers []User
	json.Unmarshal(subscribersJson, &subscribers)
	fmt.Println(subscribers)

	m := gomail.NewMessage()
	m.SetHeader("From", user.Email)
	var subscribersEmail []string
	for _, subscriber := range subscribers {
		subscribersEmail = append(subscribersEmail, subscriber.Email)
	}
	m.SetHeader("To", subscribersEmail...)

	m.SetHeader("Subject", "Today's Kakao")
	m.SetBody("text/html", contents.String())

	d := gomail.NewDialer("smtp.kakao.com", 465, user.Email, user.Password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

}
