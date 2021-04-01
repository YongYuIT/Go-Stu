package handler

import (
	"github.com/gocolly/colly"
	"strings"
	"thinking_spider/spider_interface"
)

func GetRegionHandler(spider *spider_interface.Spider) colly.HTMLCallback {
	return func(element *colly.HTMLElement) {
		region_name := strings.TrimSpace(element.Text)
		spider.SetPageValue(element.Request.URL.String(), spider_interface.REGION_NAME, region_name)
	}
}
