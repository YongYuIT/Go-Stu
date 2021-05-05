package new_release

import (
	"fmt"
	"github.com/gocolly/colly"
	"net/url"
	"strings"
	"testing"
	"thinking_spider/config"
	"thinking_spider/spider_interface"
	"thinking_spider/utils"
)

func Test_get_tree(test *testing.T) {
	homeAndKitchen := make(map[string]interface{})
	config.InitHomeAndKitchen(homeAndKitchen)
	printMap(homeAndKitchen, "")
}

func printMap(mapData map[string]interface{}, startSpace string) {
	for key := range mapData {
		value := mapData[key].(map[string]interface{})
		fmt.Println(startSpace+"-->", key)
		if len(value) > 0 {
			printMap(value, "    "+startSpace)
		}
	}
}

func Test_Tree(test *testing.T) {

	homeAndKitchen := make(map[string]interface{})
	config.InitHomeAndKitchen(homeAndKitchen)

	newReleaseTypesSpider := spider_interface.NewSpider()
	//newReleaseTypesSpider.Config.MaxDeep = 3 这里不要
	newReleaseTypesSpider.Ctrl.OnHTML("ul#zg_browseRoot", func(element *colly.HTMLElement) {

		lastPage := utils.GetUrlValueByKey(element.Request.URL.String(), "last_page")

		currentSelect := ""
		element.ForEach("li span[class='zg_selected']", func(i int, element *colly.HTMLElement) {
			currentSelect = strings.TrimSpace(element.Text)
			fmt.Println("current select-->", currentSelect)
		})
		subTree := newReleaseTypesSpider.GetPageValue(element.Request.URL.String(), "tree").(map[string]interface{})
		if subTree == nil || len(subTree) < 1 {
			fmt.Println("visit road-->", lastPage)
			fmt.Println("get page data start with-->", currentSelect, "-->", element.Request.URL.String())
		} else {
			element.ForEach("li[class!='zg_browseUp'] a[href]", func(i int, element *colly.HTMLElement) {
				tagName := strings.TrimSpace(element.Text)
				tagUrl := element.Attr("href")
				for key := range subTree {
					if strings.EqualFold(strings.TrimSpace(key), tagName) {
						tagUrl += "?last_page=" + url.QueryEscape(lastPage+"##"+currentSelect)
						newReleaseTypesSpider.SetPageValue(tagUrl, "tree", subTree[key])
						element.Request.Visit(tagUrl)
					}
				}
			})
		}

	})
	newReleaseTypesSpider.BuildStartUrl(func(spiderConfig *config.SpiderConfig) string {
		startUrl := "https://www.amazon.com/gp/new-releases/kitchen/"
		newReleaseTypesSpider.Ctrl.SetCookies(startUrl, spiderConfig.Cookies)
		newReleaseTypesSpider.SetPageValue(startUrl, "tree", homeAndKitchen["Kitchen & Dining"])
		return startUrl
	})
	newReleaseTypesSpider.StartSpider()
}
