package spider

import (
	"thinking_spider/handler"
	"thinking_spider/spider_interface"
)

func GetPatentSpider() *spider_interface.Spider {
	mainspider := spider_interface.NewSpider()
	mainspider.Config.MaxDeep = 10000
	mainspider.Ctrl.OnHTML(handler.GetPatentsListItemHandler(mainspider))
	mainspider.Ctrl.OnHTML(handler.GetPatentsPageItemHandler(mainspider))
	return mainspider
}
