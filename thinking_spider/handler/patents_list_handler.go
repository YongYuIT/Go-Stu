package handler

import (
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
	"thinking_spider/config"
	"thinking_spider/model"
	"thinking_spider/spider_interface"
)

func GetPatentsListItemHandler(thisSpider *spider_interface.Spider) (string, colly.HTMLCallback) {
	callback := func(element1 *colly.HTMLElement) {
		tableFilter := element1.ChildText("TH[scope='col']")
		if strings.Contains(tableFilter, "PAT. NO.") || strings.Contains(tableFilter, "PUB. APP. NO.") {
			element1.ForEach("TR", func(i int, element2 *colly.HTMLElement) {
				if strings.Contains(element2.ChildAttr("TD", "valign"), "top") {
					fmt.Println("------------------start")
					record := &model.PationsRecord{}
					if strings.Contains(element2.Request.URL.String(), "patft.uspto.gov") {
						record.Status = "patft"
					} else if strings.Contains(element2.Request.URL.String(), "appft.uspto.gov") {
						record.Status = "appft"
					}
					element2.ForEach("TD", func(i int, element *colly.HTMLElement) {
						if i == 0 {
							fmt.Println("index-->", element.Text)
							index, err := strconv.Atoi(strings.TrimSpace(element.Text))
							if err == nil {
								record.Index = index
							}
						}
						if i == 1 {
							fmt.Println("p.id-->", element.Text)
							record.PID = element.Text
						}
						if (i == 3 && strings.EqualFold(record.Status, "patft")) || (i == 2 && strings.EqualFold(record.Status, "appft")) {
							fmt.Println("title-->", strings.TrimSpace(element.Text))
							record.Title = strings.TrimSpace(element.Text)
							detailUrl := thisSpider.Config.WebSite + element.ChildAttr("a", "href")
							record.DetailUrl = detailUrl
							fmt.Println("url-->", detailUrl)
							spider := GetPatengsDetailSpider()
							spider.BuildStartUrl(func(spiderConfig *config.SpiderConfig) string {
								return detailUrl
							})
							fmt.Println("set record to-->", detailUrl)
							spider.SetPageValue("ALL", "record", record)
							spider.StartSpider()
						}
					})
					fmt.Println("------------------end")
				}
			})
		}
	}
	return "TABLE", callback
}
