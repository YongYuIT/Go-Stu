package spider

import "thinking_spider/spider_interface"

func GetDetailDataSpider() *spider_interface.Spider {
	detailSpider := spider_interface.NewSpider()
	detailSpider.Ctrl.OnHTML("", nil)
	return detailSpider
}
