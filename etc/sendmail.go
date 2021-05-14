package etc

import (
	"gopkg.in/gomail.v2"
	"strconv"
)

func SendMail(contents string) {

	sender := GetSenders()
	subscribers := GetSubscribers()

	m := gomail.NewMessage()
	m.SetHeader("From", sender.Email)
	m.SetBody("text/html", contents)
	intPort, err := strconv.Atoi(sender.Port)
	CheckErr(err)

	d := gomail.NewDialer(sender.Host, intPort, sender.Email, sender.Password)
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
