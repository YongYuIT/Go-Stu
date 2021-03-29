package handler

import (
	"github.com/gocolly/colly"
	"strings"
	"thinking_spider/config"
	"thinking_spider/utils"
)

func PageHandler(element *colly.HTMLElement) {
	currentSelectedPageUrl := element.ChildAttr(config.CurrentSprierConfig.PageCurrentQue, config.CurrentSprierConfig.PageAttr)
	currentPageInfo := utils.GetPageInfo(config.CurrentSprierConfig.WebSite+currentSelectedPageUrl, config.CurrentSprierConfig.PagesKey)
	element.ForEach(config.CurrentSprierConfig.PageNextQue, func(i int, pageListItem *colly.HTMLElement) {
		nextPageUrl := pageListItem.Attr(config.CurrentSprierConfig.PageAttr)
		nextTag := utils.GetNextPageStr(currentPageInfo)
		if strings.Contains(nextPageUrl, nextTag) {
			pageListItem.Request.Visit(config.CurrentSprierConfig.WebSite + nextPageUrl)
		}
	})
}
