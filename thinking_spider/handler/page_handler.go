package handler

import (
	"github.com/gocolly/colly"
	"strings"
	"thinking_spider/config"
	"thinking_spider/utils"
)

func GetPageHandler(config *config.SpiderConfig) colly.HTMLCallback {
	return func(element *colly.HTMLElement) {
		currentSelectedPageUrl := element.ChildAttr(config.PageCurrentQue, config.PageAttr)
		currentPageInfo := utils.GetUrlValueByKey(config.WebSite+currentSelectedPageUrl, config.PagesKey)
		element.ForEach(config.PageNextQue, func(i int, pageListItem *colly.HTMLElement) {
			nextPageUrl := pageListItem.Attr(config.PageAttr)
			nextTag := utils.GetNextPageStr(currentPageInfo)
			if strings.Contains(nextPageUrl, nextTag) {
				pageListItem.Request.Visit(config.WebSite + nextPageUrl)
			}
		})
	}
}
