package new_release

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
	"thinking_spider/config"
	"thinking_spider/spider"
	"thinking_spider/utils"
	"time"
)

func Test_hello(test *testing.T) {
	c := colly.NewCollector(
		colly.MaxDepth(1),
		func(collector *colly.Collector) {
			extensions.RandomUserAgent(collector)
			extensions.Referer(collector)
		})
	c.MaxDepth = 3
	c.SetRequestTimeout(time.Second * 60)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
		//降低爬取频率
		time.Sleep(2 * time.Second)
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Println("resp-code-->", response.Request.URL, "-->", response.StatusCode)
		logPath := "./logs/tmp_html/"
		os.MkdirAll(logPath, os.ModePerm)
		file, err := os.Create(fmt.Sprintf(logPath+"%d.html", time.Now().Unix()))
		if err != nil {
			return
		}
		defer file.Close()
		fmt.Fprintf(file, "<!-- "+response.Request.URL.String()+" -->\n")
		file.Write(response.Body)
	})

	c.OnHTML("ul#zg_browseRoot", func(element *colly.HTMLElement) {
		lastPage := utils.GetUrlValueByKey(element.Request.URL.String(), "last_page")
		fmt.Println("start anys page-----------------------start")
		currentSelect := ""
		element.ForEach("li span[class='zg_selected']", func(i int, element *colly.HTMLElement) {
			currentSelect = strings.TrimSpace(element.Text)
			fmt.Println("current select-->", currentSelect)
		})
		element.ForEach("li[class!='zg_browseUp'] a[href]", func(i int, element *colly.HTMLElement) {
			tagName := element.Text
			tagUrl := element.Attr("href")
			fmt.Println(lastPage, "-->", currentSelect, ": get link-->", tagName, "-->", tagUrl)
			addAttr := url.QueryEscape(lastPage + "##" + currentSelect)
			nextPage := element.Attr("href") + "?last_page=" + addAttr
			if strings.Count(lastPage, "##") >= 2 {
				//start spider
				pageSpider := spider.GetPageDataSpiderByNewRease()
				//内容页
				pageSpider.BuildStartUrl(func(spiderConfig *config.SpiderConfig) string {
					pageSpider.Ctrl.SetCookies(tagUrl, c.Cookies(element.Request.URL.String()))
					return tagUrl
				})
				pageSpider.Config.KeyWords = lastPage + "##" + currentSelect + "##" + tagName
				pageSpider.StartSpider()

			} else {
				cookie := c.Cookies(element.Request.URL.String())
				c.SetCookies(nextPage, cookie)
				element.Request.Visit(nextPage)
			}
		})
		fmt.Println("start anys page-----------------------end")
	})

	sessionId := &http.Cookie{}
	sessionId.Name = "session-id"
	sessionId.Value = "147-1594123-2406033"
	sessionId.Domain = ".amazon.com"
	sessionId.Path = "/"

	ubidMain := &http.Cookie{}
	ubidMain.Name = "ubid-main"
	ubidMain.Value = "134-3831618-4351141"
	ubidMain.Domain = ".amazon.com"
	ubidMain.Path = "/"

	cookies := []*http.Cookie{}
	cookies = append(cookies, sessionId, ubidMain)

	startUrl := "https://www.amazon.com/gp/new-releases/kitchen/"
	c.SetCookies(startUrl, cookies)

	c.Visit(startUrl)
}

func Test_getPageNum(test *testing.T) {
	num := utils.GetPageNum("Fuck111")
	fmt.Println(num)
	nextPage := utils.GetNextPageStr("Fuck9")
	fmt.Println(nextPage)
}
