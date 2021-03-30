package spider

import (
	"thinking_spider/handler"
	"thinking_spider/spider_interface"
)

func GetPagesDataSpider() *spider_interface.Spider {

	spider := spider_interface.NewSpider()
	spider.Ctrl.OnHTML(spider.Config.PageHandlerQue, handler.GetPageHandler(spider.Config))
	spider.Ctrl.OnHTML(spider.Config.ProductItemsHandlerQue, handler.GetProductItemHandler(spider.Config))

	return spider
}
