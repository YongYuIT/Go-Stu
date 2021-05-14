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
	pageVals map[string]interface{}
}

type PageKey struct {
	urlString string
	keyName   string
}

func (this *PageKey) string() string {
	return this.urlString + "##$$##&&##" + this.keyName
}

func (this *Spider) GetPageValue(url string, name string) interface{} {
	pageKey := &PageKey{
		urlString: url,
		keyName:   name,
	}
	return this.pageVals[pageKey.string()]
}

func (this *Spider) SetPageValue(url string, name string, value interface{}) {
	pageKey := &PageKey{
		urlString: url,
		keyName:   name,
	}
	this.pageVals[pageKey.string()] = value
}

func (this *Spider) SetHandler(goquerySelector string, f colly.HTMLCallback) {
	this.Ctrl.OnHTML(goquerySelector, f)
}

const (
	DEBUG_MODEL = "debug"
	REGION_NAME = "region_name"
)

func NewSpider() *Spider {

	spider := &Spider{}

	cpyConfig := (*config.CurrentDefaultConfig)
	spider.Config = &cpyConfig
	spider.pageVals = make(map[string]interface{})

	spider.Ctrl = colly.NewCollector(
		colly.MaxDepth(spider.Config.MaxDeep),
		func(collector *colly.Collector) {
			extensions.RandomUserAgent(collector)
			extensions.Referer(collector)
		},
	)
	//spider.Ctrl.SetRequestTimeout(time.Second * 60)
	spider.Ctrl.SetRequestTimeout(time.Second * time.Duration(spider.Config.TimeOut))
	fmt.Println("timeout-->", spider.Config.TimeOut)
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
			file.Write(response.Body)
		}
	})

	spider.Ctrl.OnError(func(response *colly.Response, err error) {
		fmt.Println("find a error when req-->",
			response.Request.URL.String(),
			"\n time:", time.Now().Format(time.RFC850),
			"\n keyword: ", spider.Config.KeyWords,
			"\n details for error", err)
	})

	spider.Ctrl.OnHTML(spider.Config.RobortQue, func(element *colly.HTMLElement) {
		if strings.Contains(element.Text, "robot") {
			cookie := spider.Ctrl.Cookies(element.Request.URL.String())
			fmt.Println("ant robot when req-->",
				element.Request.URL.String(),
				"\n time:", time.Now().Format(time.RFC850),
				"\n keyword: ", spider.Config.KeyWords)
			if len(cookie) > 0 {
				for i := range cookie {
					fmt.Println("cookies: ", cookie[i].Name, "-->", cookie[i].Value)
				}
			} else {
				fmt.Println("no cookies")
			}
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
