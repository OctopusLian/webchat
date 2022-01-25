package utils

import (
	"github.com/astaxie/beego"
	"gopkg.in/gomail.v2"
)

func SendEmail(to, subject, body string) error {
	smtphost := beego.AppConfig.String("smtphost")
	smtpport, _ := beego.AppConfig.Int("smtpport")
	smtpuser := beego.AppConfig.String("smtpuser")
	smtppassword := beego.AppConfig.String("smtppassword")

	m := gomail.NewMessage()
	m.SetHeader("From", smtpuser)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	d := gomail.NewDialer(smtphost, smtpport, smtpuser, smtppassword)
	return d.DialAndSend(m)
}
