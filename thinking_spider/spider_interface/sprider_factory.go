package spider_interface

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"thinking_spider/config"
	"time"
)

type Spider struct {
	Ctrl     *colly.Collector
	startUrl string
	Config   *config.SpiderConfig
}

func NewSpider() *Spider {

	spider := &Spider{}

	cpyConfig := (*config.CurrentDefaultConfig)
	spider.Config = &cpyConfig

	spider.Ctrl = colly.NewCollector(
		colly.MaxDepth(spider.Config.MaxDeep),
		func(collector *colly.Collector) {
			extensions.RandomUserAgent(collector)
			extensions.Referer(collector)
		},
	)

	spider.Ctrl.OnResponse(func(response *colly.Response) {
		fmt.Println("resp-code-->", response.Request.URL, "-->", response.StatusCode)
	})

	return spider
}

func (this *Spider) BuildStartUrl(builder func(spiderConfig *config.SpiderConfig) string) {
	this.startUrl = builder(this.Config)
}

func (this *Spider) StartSpider() error {

	this.Ctrl.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		//降低爬取频率
		time.Sleep(time.Duration(this.Config.DelaySpider) * time.Second)
	})

	this.Ctrl.MaxDepth = this.Config.MaxDeep

	return this.Ctrl.Visit(this.startUrl)
}
