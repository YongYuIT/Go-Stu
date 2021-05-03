package spider

import (
	"thinking_spider/handler"
	"thinking_spider/spider_interface"
)

func GetPriceLevelDataSpider() *spider_interface.Spider {
	priceLevelspider := spider_interface.NewSpider()
	priceLevelHandler := handler.GetPriceLevelHandler(priceLevelspider, GetPagesDataSpiderByKeyWord())
	//priceLevelspider.Ctrl.OnHTML(priceLevelspider.Config.PriceLevelConfig.PriceListQue, priceLevelHandler)
	priceLevelspider.Ctrl.OnHTML("div[class='a-section a-spacing-double-large']", priceLevelHandler)
	return priceLevelspider
}
