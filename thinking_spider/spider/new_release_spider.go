package spider

import (
	"thinking_spider/handler"
	"thinking_spider/spider_interface"
)

func GetNewReleaseTypeSpider() *spider_interface.Spider {
	newReleaseTypesSpider := spider_interface.NewSpider()
	newReleaseTypesHandler := handler.GetNewReleaseTypesHandler(newReleaseTypesSpider, GetPageDataSpiderByNewRease())
	newReleaseTypesSpider.Ctrl.OnHTML("ul#zg_browseRoot", newReleaseTypesHandler)
	return newReleaseTypesSpider
}
