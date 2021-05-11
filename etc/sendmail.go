package etc

import (
	"encoding/json"
	"fmt"
	"gopkg.in/gomail.v2"
	"io/ioutil"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func SendMail(contents string) {
	userJson, err := ioutil.ReadFile("./secrets/sendmail.json") // articles.json 파일의 내용을 읽어서 바이트 슬라이스에 저장
	CheckErr(err)
	var user User
	json.Unmarshal(userJson, &user)

	subscribersJson, err := ioutil.ReadFile("./secrets/subscribers.json") // articles.json 파일의 내용을 읽어서 바이트 슬라이스에 저장
	CheckErr(err)
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
	m.SetBody("text/html", contents)

	d := gomail.NewDialer("smtp.kakao.com", 465, user.Email, user.Password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
