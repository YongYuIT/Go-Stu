package new_release

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"net/http"
	"os"
	"strings"
	"testing"
	"thinking_spider/config"
	_ "thinking_spider/config"
	"thinking_spider/spider"
	"thinking_spider/utils"
	"time"
)

func Test_bug_fix(test *testing.T) {
	newReleaseTypesSpider := spider.GetNewReleaseTypeSpider()
	newReleaseTypesSpider.Config.MaxDeep = 3
	newReleaseTypesSpider.BuildStartUrl(func(spiderConfig *config.SpiderConfig) string {
		startUrl := newReleaseTypesSpider.Config.NewRelease
		newReleaseTypesSpider.Ctrl.SetCookies(startUrl, spiderConfig.Cookies)
		return startUrl
	})
	newReleaseTypesSpider.StartSpider()
}

func Test_get_level(test *testing.T) {
	level := utils.GetLevelOfUrl("https://www.amazon.com/gp/new-releases/books/3564986011/ref=zg_bsnr_nav_b_4_266162")
	fmt.Println("level-->", level)
	level = utils.GetLevelOfUrl("https://www.amazon.com/gp/new-releases/home-garden/289668/ref=zg_bsnr_unv_k_2_289669_1")
	fmt.Println("level-->", level)
}

func Test_find_next_page(test *testing.T) {
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

	c.OnHTML("ul#zg_browseRoot", func(element1 *colly.HTMLElement) {
		currentSelect := strings.TrimSpace(element1.ChildText("li span[class='zg_selected']"))
		currentLevel := utils.GetLevelOfUrl(element1.Request.URL.String())
		sameLevels := []string{}
		element1.ForEach("li span[class='zg_selected']", func(i int, element *colly.HTMLElement) {
			node := element.DOM.Parent().Parent()
			node.Find("li").Each(func(i int, selection *goquery.Selection) {
				fmt.Println("get li the same level-->", selection.Text())
				sameLevels = append(sameLevels, strings.TrimSpace(selection.Text()))
			})
		})

		hasNext := false
		element1.ForEach("li[class!='zg_browseUp'] a[href]", func(i int, element *colly.HTMLElement) {
			tagUrl := element.Attr("href")
			tagLevel := utils.GetLevelOfUrl(tagUrl)
			tagName := element.Text
			if tagLevel > currentLevel && !utils.StrInList(tagName, sameLevels) {
				hasNext = true
			}
		})
		fmt.Println("current-->", currentSelect, "-->hasNext-->", hasNext)
		if hasNext {
			element1.ForEach("li[class!='zg_browseUp'] a[href]", func(i int, element *colly.HTMLElement) {
				tagUrl := element.Attr("href")
				tagName := element.Text
				tagLevel := utils.GetLevelOfUrl(tagUrl)
				if tagLevel >= currentLevel {
					fmt.Println("get next page-->", tagName)
				}
			})
		} else {
			fmt.Println("read the page")
		}
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

	startUrl := "https://www.amazon.com/gp/new-releases/kitchen/289671/ref=zg_bsnr_unv_k_3_289815_1"
	//startUrl := "https://www.amazon.com/gp/new-releases/kitchen/3737261/ref=zg_bsnr_nav_k_3_289671"
	c.SetCookies(startUrl, cookies)

	c.Visit(startUrl)
}
