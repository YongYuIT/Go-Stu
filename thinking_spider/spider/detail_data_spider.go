package spider

import (
	"thinking_spider/model"
	"thinking_spider/spider_interface"
)
import "thinking_spider/handler"

func GetDetailDataSpider() *spider_interface.Spider {
	detailSpider := spider_interface.NewSpider()
	record := &model.ProdDetailRecord{}
	detailSpider.Ctrl.OnHTML(handler.Get5DescsHandler(detailSpider, record))
	detailSpider.Ctrl.OnHTML(handler.GetAsin(detailSpider, record))
	detailSpider.Ctrl.OnHTML(handler.GetSoldBy(detailSpider, record))
	detailSpider.Ctrl.OnHTML(handler.GetProdDesc(detailSpider, record))
	detailSpider.Ctrl.OnHTML(handler.GetSoldID(detailSpider, record))
	detailSpider.Ctrl.OnHTML(handler.GetDataFirstAvailable(detailSpider, record))
	return detailSpider
}
