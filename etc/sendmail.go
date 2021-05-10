package etc

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

func SendMail(jobs []byte) {
	m := gomail.NewMessage()
	m.SetHeader("From", "bogus92@kakao.com")
	m.SetHeader("To", "bogus92@kakao.com")
	m.SetAddressHeader("Cc", "bogus92@kakao.com", "Dan")
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "Hello <b>Bob</b> and <i>Cora</i>!")

	fmt.Println(string(jobs))

	d := gomail.NewDialer("smtp.kakao.com", 465, "bogus92@kakao.com", "zes200101")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
