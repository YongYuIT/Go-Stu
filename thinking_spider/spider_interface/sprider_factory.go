package spider_interface

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"os"
	"strings"
	"thinking_spider/config"
	"time"
)

type Spider struct {
	Ctrl     *colly.Collector
	startUrl string
	Config   *config.SpiderConfig
}

const (
	DEBUG_MODEL = "debug"
)

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
		if strings.EqualFold(DEBUG_MODEL, spider.Config.Model) {
			os.MkdirAll("./logs/tmp_html/", os.ModePerm)
			file, err := os.Create(fmt.Sprintf("./logs/tmp_html/%d.html", time.Now().Unix()))
			if err != nil {
				return
			}
			defer file.Close()
			fmt.Fprintf(file, "<!-- "+response.Request.URL.String()+" -->\n")
			fmt.Fprintf(file, string(response.Body))
		}
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
