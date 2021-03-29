package test

import (
	"fmt"
	"github.com/gocolly/colly"
	"testing"
	"thinking_spider/config"
	_ "thinking_spider/config"
	_ "thinking_spider/database"
	"thinking_spider/handler"
	"thinking_spider/utils"
)

func Test_pages_data(test *testing.T) {
	startUrl := utils.GetUrl(config.CurrentSprierConfig.WebSite, []string{"k=" + utils.GetKeyWords(config.CurrentSprierConfig.KeyWords), "ref=nb_sb_noss"})
	fmt.Println("startUrl-->", startUrl)

	collyPro := colly.NewCollector(
		colly.MaxDepth(config.CurrentSprierConfig.MaxDeep),
	)
	collyPro.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting-->", r.URL)
	})

	collyPro.OnResponse(func(response *colly.Response) {
		fmt.Println("resp-code-->", response.Request.URL, "-->", response.StatusCode)
	})
	collyPro.OnHTML(config.CurrentSprierConfig.PageHandlerQue, handler.PageHandler)
	collyPro.OnHTML(config.CurrentSprierConfig.ProductItemsHandlerQue, handler.ProductItemHandler)

	collyPro.Visit(startUrl)

}
