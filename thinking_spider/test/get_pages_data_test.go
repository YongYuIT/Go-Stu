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
	startUrl := utils.GetUrlWithKVs(config.CurrentDefaultConfig.WebSite, []string{"k=" + utils.GetKeyWords(config.CurrentDefaultConfig.KeyWords), "ref=nb_sb_noss"})
	fmt.Println("startUrl-->", startUrl)

	collyPro := colly.NewCollector(
		colly.MaxDepth(config.CurrentDefaultConfig.MaxDeep),
	)
	collyPro.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting-->", r.URL)
	})

	collyPro.OnResponse(func(response *colly.Response) {
		fmt.Println("resp-code-->", response.Request.URL, "-->", response.StatusCode)
	})
	collyPro.OnHTML(config.CurrentDefaultConfig.PageHandlerQue, handler.GetPageHandler(config.CurrentDefaultConfig))
	collyPro.OnHTML(config.CurrentDefaultConfig.ProductItemsHandlerQue, handler.GetProductItemHandler(config.CurrentDefaultConfig))

	collyPro.Visit(startUrl)

}

func Test_CNY(test *testing.T) {
	fmt.Println(utils.GetPrice("¥71.93"))
}
