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
	pageVals map[string]string
}

type PageKey struct {
	urlString string
	keyName   string
}

func (this *PageKey) string() string {
	return this.urlString + "##$$##&&##" + this.keyName
}

func (this *Spider) GetPageValue(url string, name string) string {
	pageKey := &PageKey{
		urlString: url,
		keyName:   name,
	}
	return this.pageVals[pageKey.string()]
}

func (this *Spider) SetPageValue(url string, name string, value string) {
	pageKey := &PageKey{
		urlString: url,
		keyName:   name,
	}
	this.pageVals[pageKey.string()] = value
}

const (
	DEBUG_MODEL = "debug"
	REGION_NAME = "region_name"
)

func NewSpider() *Spider {

	spider := &Spider{}

	cpyConfig := (*config.CurrentDefaultConfig)
	spider.Config = &cpyConfig
	spider.pageVals = make(map[string]string)

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
			logPath := "./logs/tmp_html/"
			os.MkdirAll(logPath, os.ModePerm)
			file, err := os.Create(fmt.Sprintf(logPath+"%d.html", time.Now().Unix()))
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
