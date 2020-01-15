package tools

import (
	"../model"
	"gopkg.in/gomail.v2"
	"strconv"
)

type EmailTool struct {
}

func (thiz *EmailTool) sendEmail(email *model.Email) bool {
	conf := GetSendEmailConf()
	port, err := strconv.Atoi(conf.STMPPort)
	if err != nil {
		return false
	}
	d := gomail.NewDialer(conf.STMPHost, port, conf.Username, conf.Passwd)
	err = d.DialAndSend(convertEmailModel2EmailMsg(email))
	if err != nil {
		return false
	}
	return true
}

func convertEmailModel2EmailMsg(email *model.Email) *gomail.Message {
	msg := gomail.NewMessage()
	msg.SetHeader("From", msg.FormatAddress(email.From, "生产环境监控服务"))
	msg.SetHeader("To", email.To)
	msg.SetHeader("Cc", email.Cc)
	msg.SetHeader("Bbcc", email.Bcc)
	msg.SetHeader("Subject", email.Title)
	msg.SetBody("text/html", email.Content)
	return msg
}
