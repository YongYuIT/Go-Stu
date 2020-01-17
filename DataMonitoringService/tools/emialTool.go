package tools

import (
	"../model"
	"fmt"
	"gopkg.in/gomail.v2"
	"strconv"
)

type EmailTool struct {
}

func (thiz *EmailTool) SendEmail(email *model.Email) bool {
	conf := GetSendEmailConf()
	port, err := strconv.Atoi(conf.STMPPort)
	if err != nil {
		return false
	}
	d := gomail.NewDialer(conf.STMPHost, port, conf.Username, conf.Passwd)
	email.From = conf.Username
	err = d.DialAndSend(convertEmailModel2EmailMsg(email))
	if err != nil {
		fmt.Println("email err", err)
		return false
	}
	return true
}

func convertEmailModel2EmailMsg(email *model.Email) *gomail.Message {
	msg := gomail.NewMessage()
	msg.SetHeader("From", msg.FormatAddress(email.From, "生产环境监控服务"))
	msg.SetHeader("To", email.To)
	email.Cc = email.From
	msg.SetHeader("Cc", email.Cc)

	if email.Bcc != "" {
		msg.SetHeader("Bcc", email.Bcc)
	}
	msg.SetHeader("Subject", email.Title)
	msg.SetBody("text/html", email.Content)
	msg.Attach(email.Attachments)
	fmt.Println("sending email-->", email.Content)
	return msg
}
