package handler

import (
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
	"thinking_spider/model"
	"thinking_spider/spider_interface"
	"thinking_spider/utils"
	"time"
)

func GetPatentsPDFHandler(thisSpider *spider_interface.Spider) (string, colly.HTMLCallback) {
	callback := func(element *colly.HTMLElement) {
		pdfUrl := element.Attr("src")[2:]
		fmt.Println("get pdf-->", pdfUrl)
		pdfPath := "./tmp_pdf/"
		pdfName := strconv.FormatInt(time.Now().Unix(), 10) + ".pdf"
		fmt.Println("save to-->", pdfName)
		fmt.Println("get record from-->", element.Request.URL.String())
		record := thisSpider.GetPageValue("ALL", "record").(*model.PationsRecord)
		record.PDFUrl = "https://" + pdfUrl
		isSuccess := utils.DownloadFile(record.PDFUrl, pdfPath, pdfName)
		if isSuccess {
			record.PDFPath = pdfPath + pdfName
			simg := utils.CreateSImgFromPatentsPDF("", record.PDFPath)
			if !strings.EqualFold(simg, "") {
				record.IMGPath = simg
			}
		}
		model.SavePationsRecord(record)
	}
	return "embed[type='application/pdf']", callback
}
