package handler

import (
	"thinking_spider/spider_interface"
)

func GetPatengsDetailSpider() *spider_interface.Spider {
	spider := spider_interface.NewSpider()
	spider.Config.MaxDeep = 2
	spider.Ctrl.OnHTML(GetPatentsPDFHandler(spider))
	spider.Ctrl.OnHTML(GetPatentsIMGHandler(spider))
	return spider
}
