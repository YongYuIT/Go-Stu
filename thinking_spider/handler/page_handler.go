package handler

import (
	"github.com/gocolly/colly"
	"strings"
	"thinking_spider/spider_interface"
	"thinking_spider/utils"
)

func GetPageHandler(pageSpider *spider_interface.Spider) colly.HTMLCallback {
	config := pageSpider.Config
	return func(element *colly.HTMLElement) {
		currentSelectedPageUrl := element.ChildAttr(config.PageCurrentQue, config.PageAttr)
		currentPageInfo := utils.GetUrlValueByKey(config.WebSite+currentSelectedPageUrl, config.PagesKey)
		element.ForEach(config.PageNextQue, func(i int, pageListItem *colly.HTMLElement) {
			nextPageUrl := pageListItem.Attr(config.PageAttr)
			nextTag := utils.GetNextPageStr(currentPageInfo)
			if strings.Contains(nextPageUrl, nextTag) {
				cookie := pageSpider.Ctrl.Cookies(element.Request.URL.String())
				nextPageUrl = config.WebSite + nextPageUrl
				pageSpider.Ctrl.SetCookies(nextPageUrl, cookie)
				pageListItem.Request.Visit(nextPageUrl)
			}
		})
	}
}
