package price_test

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"net/url"
	"strings"
	"testing"
	"thinking_spider/spider"
	"time"
)

const website = "https://www.amazon.com"

func Test_go_to_next_page_price(test *testing.T) {

	c := colly.NewCollector(
		colly.MaxDepth(2),
		func(collector *colly.Collector) {
			extensions.RandomUserAgent(collector)
			extensions.Referer(collector)
		},
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		//降低爬取频率
		time.Sleep(time.Duration(2) * time.Second)
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Println("resp-code-->", response.Request.URL, "-->", response.StatusCode)
	})

	c.OnHTML("div#priceRefinements a[data-routing][class='a-link-normal s-navigation-item']", func(element *colly.HTMLElement) {
		thkPTag := GetUrlValueByKey(element.Request.URL.String(), "thk_p_tag")

		if strings.EqualFold(thkPTag, "") {
			//起始页

			priceUrl := element.Attr("href")
			priceStr := element.ChildText("span[class='a-size-base a-color-base']")
			fmt.Println(priceUrl, "-->", priceStr)

			cookie := c.Cookies(element.Request.URL.String())
			priceUrl = website + priceUrl + "&thk_p_tag=" + url.QueryEscape(priceStr)
			c.SetCookies(priceUrl, cookie)
			element.Request.Visit(priceUrl)

		} else {
			//内容页
			startUrl := element.Request.URL.String()
			spider := spider.GetPagesDataSpiderByKeyWord(startUrl)
			spider.Config.WebSite = website
			spider.StartSpider()
		}
	})
	c.Visit(website + "/s?k=coat+hangers&ref=nb_sb_noss")
}

func GetUrlValueByKey(urlstr string, key string) string {
	thisUrl, err := url.Parse(urlstr)
	if err != nil {
		return ""
	}
	kvs, err := url.ParseQuery(thisUrl.RawQuery)
	if err != nil {
		return ""
	}
	if len(kvs[key]) > 0 {
		return kvs[key][0]
	} else {
		return ""
	}
}
