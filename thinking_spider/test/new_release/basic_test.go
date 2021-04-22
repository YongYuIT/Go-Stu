package new_release

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"net/http"
	"os"
	"testing"
	"thinking_spider/handler"
	"thinking_spider/spider"
	"time"
)

func Test_basic(test *testing.T) {
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

	spider := spider.GetBasicPageDataSpider()
	spider.Config.KeyWords = "Kitchen & Dining##Bakeware##Bakers & Casseroles##Casseroles"
	c.OnHTML("ol#zg-ordered-list", handler.GetNewReleaseItemHandler(spider))

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

	startUrl := "https://www.amazon.com/gp/new-releases/kitchen/289815/ref=zg_bsnr_nav_k_3_289671"
	c.SetCookies(startUrl, cookies)

	c.Visit(startUrl)
}
