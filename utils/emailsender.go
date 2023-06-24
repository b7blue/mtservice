package utils

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

func SendEmail(email, code string) error {
	// 发邮件提醒
	m := gomail.NewMessage()
	m.SetHeader("From", "mtslash_service@163.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "mtservice注册验证码")

	m.SetBody("text/plain", fmt.Sprintf("本次邮箱注册验证码为:%s\n验证码有效时间为10分钟,请尽快填写:)", code))

	d := gomail.NewDialer("smtp.163.com", 465, "mtslash_service@163.com", "VDLCWZSSZIXWLHCW")

	return d.DialAndSend(m)
}
