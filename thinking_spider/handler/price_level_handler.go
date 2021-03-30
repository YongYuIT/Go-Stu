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

const thkPTag = "thk_p_tag"

func GetPriceLevelHandler(priceLevelSpider *spider_interface.Spider, pageSpider *spider_interface.Spider) colly.HTMLCallback {
	return func(element *colly.HTMLElement) {
		currentthkPTag := utils.GetUrlValueByKey(element.Request.URL.String(), thkPTag)
		if strings.EqualFold(currentthkPTag, "") {
			//起始页
			priceUrl := element.Attr(priceLevelSpider.Config.PageAttr)
			priceStr := element.ChildText(priceLevelSpider.Config.PriceLevelConfig.PriceStrQue)
			fmt.Println(priceUrl, "-->", priceStr)

			cookie := priceLevelSpider.Ctrl.Cookies(element.Request.URL.String())
			priceUrl = priceLevelSpider.Config.WebSite + priceUrl + "&" + thkPTag + "=" + url.QueryEscape(priceStr)
			priceLevelSpider.Ctrl.SetCookies(priceUrl, cookie)
			element.Request.Visit(priceUrl)

		} else {
			//内容页
			pageSpider.BuildStartUrl(func(spiderConfig *config.SpiderConfig) string {
				return element.Request.URL.String()
			})
			level, err := url.QueryUnescape(currentthkPTag)
			if err == nil {
				pageSpider.Config.CurrentPriceLevel = level
			}
			pageSpider.StartSpider()
		}
	}
}
