package etc

import (
	"encoding/json"
	"fmt"
	"gopkg.in/gomail.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func SendMail(contents string) {
	dir, err := filepath.Abs(filepath.Dir("../secrets/"))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("realpath : ", dir)

	// TODO : GetSenders() 로 aws RDS 에서 불러오기
	userJson, err := ioutil.ReadFile(filepath.Join(dir, "sendmail.json")) // articles.json 파일의 내용을 읽어서 바이트 슬라이스에 저장
	CheckErr(err)
	var user User
	json.Unmarshal(userJson, &user)

	subscribers := GetSubscribers()

	m := gomail.NewMessage()
	m.SetHeader("From", user.Email)
	m.SetBody("text/html", contents)
	d := gomail.NewDialer("smtp.kakao.com", 465, user.Email, user.Password)
	if _, err := d.Dial(); err != nil {
		panic(err)
	}

	for _, subscriber := range subscribers {
		m.SetHeader("To", subscriber.Email)
		subject := subscriber.Name + " 님 ! 오늘의 채용정보입니다👶"
		m.SetHeader("Subject", subject)

		if err := d.DialAndSend(m); err != nil {
			panic(err)
		}
	}

}
