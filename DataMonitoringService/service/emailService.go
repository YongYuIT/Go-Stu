package service

import (
	"../model"
	"../tools"
	"bytes"
	"fmt"
	"html/template"
)

type EmailService struct {
}

func (thiz *EmailService) SendSimpleSchaInfoEmail(scha_info string, to string) {
	linFileName := (&PrintService{}).PrintSchaPolylines(scha_info)
	lineFileStr := (&tools.FileTool{}).ReadFileTxt(linFileName)
	content := struct {
		Disc string
		Line template.HTML
	}{"fuck worninggggg11111111111111", template.HTML(lineFileStr)}
	temp, err := template.ParseFiles("../view/SimpSchaEmailView.html")
	buf := new(bytes.Buffer)
	if err == nil {
		err := temp.Execute(buf, content)
		if err != nil {
			fmt.Println("html err", err)
		}
	}

	emailModel := &model.Email{}
	emailModel.Content = buf.String()
	emailModel.Title = "fuck test"
	emailModel.To = to
	emailModel.Attachments = linFileName
	(&tools.EmailTool{}).SendEmail(emailModel)

}
