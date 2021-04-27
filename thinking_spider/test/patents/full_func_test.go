package patents

import (
	"fmt"
	"testing"
	"thinking_spider/config"
	"thinking_spider/handler"
	"thinking_spider/spider_interface"
)

func Test_full_func(test *testing.T) {
	startPage := "http://patft.uspto.gov/netacgi/nph-Parser?Sect1=PTO2&Sect2=HITOFF&p=1&u=%2Fnetahtml%2FPTO%2Fsearch-bool.html&r=0&f=S&l=50&TERM1=Soap+box&FIELD1=&co1=OR&TERM2=Soap+dish&FIELD2=&d=PTXT"
	fmt.Println("start --> ", startPage)
	mainspider := spider_interface.NewSpider()
	mainspider.Config.MaxDeep = 10000
	mainspider.Config.WebSite = "http://patft.uspto.gov"
	mainspider.Ctrl.OnHTML(handler.GetPatentsListItemHandler(mainspider))
	mainspider.Ctrl.OnHTML(handler.GetPatentsPageItemHandler(mainspider))
	mainspider.BuildStartUrl(func(spiderConfig *config.SpiderConfig) string {
		return startPage
	})
	mainspider.StartSpider()
}
