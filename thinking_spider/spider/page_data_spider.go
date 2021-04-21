package spider

import (
	"thinking_spider/handler"
	"thinking_spider/spider_interface"
)

func GetPagesDataSpiderByKeyWord() *spider_interface.Spider {
	spider := GetBasicPageDataSpider()
	spider.Ctrl.OnHTML(spider.Config.ProductItemsHandlerQue, handler.GetProductItemHandler(spider))
	return spider
}

func GetPageDataSpiderByNewRease() *spider_interface.Spider {
	spider := GetBasicPageDataSpider()
	spider.Ctrl.OnHTML("ol#zg-ordered-list", handler.GetNewReaseItemHandler(spider))
	return spider
}

func GetBasicPageDataSpider() *spider_interface.Spider {
	spider := spider_interface.NewSpider()
	spider.Ctrl.OnHTML(spider.Config.PageHandlerQue, handler.GetPageHandler(spider))
	spider.Ctrl.OnHTML(spider.Config.RegionQue, handler.GetRegionHandler(spider))
	return spider
}
