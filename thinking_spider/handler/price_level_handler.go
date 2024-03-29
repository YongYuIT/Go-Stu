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
		hasPrice := false
		if strings.EqualFold(currentthkPTag, "") {
			//起始页
			element.ForEach(priceLevelSpider.Config.PriceLevelConfig.PriceListQue, func(i int, element1 *colly.HTMLElement) {
				hasPrice = true
				priceUrl := element1.Attr(priceLevelSpider.Config.PageAttr)
				priceStr := element1.ChildText(priceLevelSpider.Config.PriceLevelConfig.PriceStrQue)
				fmt.Println(priceUrl, "-->", priceStr)

				if !(strings.EqualFold(priceLevelSpider.Config.PricesLevels, "")) {
					if !strings.Contains(priceLevelSpider.Config.PricesLevels, priceStr) {
						return
					}
				}
				cookie := priceLevelSpider.Ctrl.Cookies(element1.Request.URL.String())
				priceUrl = priceLevelSpider.Config.WebSite + priceUrl + "&" + thkPTag + "=" + url.QueryEscape(priceStr)
				priceLevelSpider.Ctrl.SetCookies(priceUrl, cookie)
				element1.Request.Visit(priceUrl)
			})
		}

		if !hasPrice || !strings.EqualFold(currentthkPTag, "") {
			//内容页
			pageSpider.BuildStartUrl(func(spiderConfig *config.SpiderConfig) string {
				startUrl := element.Request.URL.String()
				//pageSpider.Ctrl.SetCookies(startUrl, priceLevelSpider.Ctrl.Cookies(element.Request.URL.String()))
				pageSpider.Ctrl.SetCookies(startUrl, spiderConfig.GetCookie())
				return startUrl
			})
			level, err := url.QueryUnescape(currentthkPTag)
			if err == nil {
				pageSpider.Config.CurrentPriceLevel = level
			}
			pageSpider.StartSpider()
		}
	}
}
