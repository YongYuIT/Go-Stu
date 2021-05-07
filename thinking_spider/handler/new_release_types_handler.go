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

		currentSelect := ""
		element.ForEach("li span[class='zg_selected']", func(i int, element *colly.HTMLElement) {
			currentSelect = strings.TrimSpace(element.Text)
			fmt.Println("current select-->", currentSelect)
		})
		subTree := newReleaseTypesSpider.GetPageValue(element.Request.URL.String(), "tree").(map[string]interface{})
		if subTree == nil || len(subTree) < 1 {
			fmt.Println("visit road-->", lastPage)
			fmt.Println("get page data start with-->", currentSelect, "-->", element.Request.URL.String())

			someTypeSpider.BuildStartUrl(func(spiderConfig *config.SpiderConfig) string {
				//someTypeSpider.Ctrl.SetCookies(element.Request.URL.String(), newReleaseTypesSpider.Ctrl.Cookies(element.Request.URL.String()))
				someTypeSpider.Ctrl.SetCookies(element.Request.URL.String(), spiderConfig.GetCookie())
				return element.Request.URL.String()
			})
			someTypeSpider.Config.KeyWords = lastPage + "##" + currentSelect
			someTypeSpider.StartSpider()

		} else {
			element.ForEach("li[class!='zg_browseUp'] a[href]", func(i int, element *colly.HTMLElement) {
				tagName := strings.TrimSpace(element.Text)
				tagUrl := element.Attr("href")
				for key := range subTree {
					if strings.EqualFold(strings.TrimSpace(key), tagName) {
						tagUrl += "?last_page=" + url.QueryEscape(lastPage+"##"+currentSelect)
						newReleaseTypesSpider.SetPageValue(tagUrl, "tree", subTree[key])
						element.Request.Visit(tagUrl)
					}
				}
			})
		}
	}
}
