package handler

import (
	"github.com/gocolly/colly"
	"strings"
	"thinking_spider/spider_interface"
)

func GetPatentsPageItemHandler(thisSpider *spider_interface.Spider) (string, colly.HTMLCallback) {
	callback := func(element1 *colly.HTMLElement) {
		if strings.Contains(element1.ChildAttr("img[align='MIDDLE']", "src"), "nextlist") {
			url := element1.Attr("href")
			isV := thisSpider.GetPageValue(url, "isVisit")
			if isV == nil || !isV.(bool) {
				nextUrl := "http://patft.uspto.gov" + url
				thisSpider.SetPageValue(url, "isVisit", true)
				element1.Request.Visit(nextUrl)
			}
		}
	}
	return "a[href]", callback
}
