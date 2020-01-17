package service

import (
	"../tools"
	"strings"
)

type EmailService struct {
}

func (thiz *EmailService) sendSimpleSchaInfoEmail(scha_info string) {

	fileTool := &tools.FileTool{}
	pService := PrintService{}
	linFileName := pService.PrintSchaPolylines(scha_info)
	viewHtml := fileTool.ReadFileTxt("../view/SimpSchaEmailView.html")
	viewHtml = strings.Replace(viewHtml, "##$text_content$##", "fuck worning", -1)
	lineFileStr := fileTool.ReadFileTxt(linFileName)
	viewHtml = strings.Replace(viewHtml, "##$html_content$##", lineFileStr, -1)
}
