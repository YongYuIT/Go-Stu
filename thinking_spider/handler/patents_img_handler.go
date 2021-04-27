package handler

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"thinking_spider/spider_interface"
)

func GetPatentsIMGHandler(thisSpider *spider_interface.Spider) (string, colly.HTMLCallback) {
	callback := func(element *colly.HTMLElement) {
		if strings.Contains(element.ChildAttr("img", "src"), "image.gif") {
			isImaged := thisSpider.GetPageValue("ALL", "isImaged")
			fmt.Println("get isImaged-->", isImaged)
			if isImaged == nil || !isImaged.(bool) {
				imgSrc := element.Attr("href")
				fmt.Println("get img src-->", imgSrc)
				element.Request.Visit(imgSrc)
			}
			thisSpider.SetPageValue("ALL", "isImaged", true)
		}
	}
	return "a[href]", callback
}
