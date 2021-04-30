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
	return func(element1 *colly.HTMLElement) {

		lastPage := utils.GetUrlValueByKey(element1.Request.URL.String(), "last_page")

		currentSelect := strings.TrimSpace(element1.ChildText("li span[class='zg_selected']"))
		currentLevel := utils.GetLevelOfUrl(element1.Request.URL.String())
		hasNext := false
		element1.ForEach("li[class!='zg_browseUp'] a[href]", func(i int, element *colly.HTMLElement) {
			tagUrl := element.Attr("href")
			tagLevel := utils.GetLevelOfUrl(tagUrl)
			if tagLevel > currentLevel {
				hasNext = true
			}
		})
		fmt.Println("current-->", currentSelect, "-->hasNext-->", hasNext)
		if hasNext {
			element1.ForEach("li[class!='zg_browseUp'] a[href]", func(i int, element *colly.HTMLElement) {
				tagUrl := element.Attr("href")
				tagName := element.Text
				tagLevel := utils.GetLevelOfUrl(tagUrl)
				if tagLevel >= currentLevel {
					addAttr := url.QueryEscape(lastPage + "##" + currentSelect)
					orgNextUrl := element.Attr("href")
					isVisit := newReleaseTypesSpider.GetPageValue(orgNextUrl, "isVisit")
					if isVisit == nil || !isVisit.(bool) {
						nextPage := orgNextUrl + "?last_page=" + addAttr
						cookie := newReleaseTypesSpider.Ctrl.Cookies(element.Request.URL.String())
						newReleaseTypesSpider.Ctrl.SetCookies(nextPage, cookie)
						element.Request.Visit(nextPage)
						fmt.Println("get next page-->", tagName)
						newReleaseTypesSpider.SetPageValue(orgNextUrl, "isVisit", true)
					}
				}
			})
		} else {
			fmt.Println("read the page")
			someTypeSpider.BuildStartUrl(func(spiderConfig *config.SpiderConfig) string {
				someTypeSpider.Ctrl.SetCookies(element1.Request.URL.String(), newReleaseTypesSpider.Ctrl.Cookies(element1.Request.URL.String()))
				return element1.Request.URL.String()
			})
			someTypeSpider.Config.KeyWords = lastPage + "##" + currentSelect
			//someTypeSpider.StartSpider()
		}

	}
}
