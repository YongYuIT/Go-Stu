package handler

import (
	"fmt"
	"github.com/gocolly/colly"
	"net/url"
	"strings"
	"thinking_spider/config"
	"thinking_spider/spider_interface"
	"thinking_spider/utils"
)

func GetNewReleaseTypesHandler(newReleaseTypesSpider *spider_interface.Spider, someTypeSpider *spider_interface.Spider) colly.HTMLCallback {
	return func(element *colly.HTMLElement) {
		lastPage := utils.GetUrlValueByKey(element.Request.URL.String(), "last_page")
		fmt.Println("start anys page-----------------------start")
		currentSelect := ""
		element.ForEach("li span[class='zg_selected']", func(i int, element *colly.HTMLElement) {
			currentSelect = strings.TrimSpace(element.Text)
			fmt.Println("current select-->", currentSelect)
		})
		element.ForEach("li[class!='zg_browseUp'] a[href]", func(i int, element *colly.HTMLElement) {
			tagName := element.Text
			tagUrl := element.Attr("href")
			fmt.Println(lastPage, "-->", currentSelect, ": get link-->", tagName, "-->", tagUrl)
			addAttr := url.QueryEscape(lastPage + "##" + currentSelect)
			nextPage := element.Attr("href") + "?last_page=" + addAttr
			if strings.Count(lastPage, "##") >= 2 {
				//内容页
				someTypeSpider.BuildStartUrl(func(spiderConfig *config.SpiderConfig) string {
					someTypeSpider.Ctrl.SetCookies(tagUrl, newReleaseTypesSpider.Ctrl.Cookies(element.Request.URL.String()))
					return tagUrl
				})
				someTypeSpider.Config.KeyWords = lastPage + "##" + currentSelect + "##" + tagName
				someTypeSpider.StartSpider()
			} else {
				cookie := newReleaseTypesSpider.Ctrl.Cookies(element.Request.URL.String())
				newReleaseTypesSpider.Ctrl.SetCookies(nextPage, cookie)
				element.Request.Visit(nextPage)
			}
		})
		fmt.Println("start anys page-----------------------end")
	}
}
