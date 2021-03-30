package simple_test

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
	"testing"
	"thinking_spider/utils"
)

const website = "https://www.amazon.cn"

func Test_go_to_next_page(test *testing.T) {
	c := colly.NewCollector(
		colly.MaxDepth(10),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Println("resp-code-->", response.Request.URL, "-->", response.StatusCode)
	})

	c.OnHTML("div[class='a-text-center']", func(element *colly.HTMLElement) {
		url := element.ChildAttr("li[class=a-selected] a[href]", "href")
		pageInfo := utils.GetUrlValueByKey(website+url, "ref")
		element.ForEach("li[class=a-normal] a[href]", func(i int, pageListItem *colly.HTMLElement) {
			nextPageUrl := pageListItem.Attr("href")
			nextTag := utils.GetNextPageStr(pageInfo)
			if strings.Contains(nextPageUrl, nextTag) {
				pageListItem.Request.Visit(website + nextPageUrl)
			}
		})
	})
	c.Visit(website + "/s?k=programming+books&__mk_zh_CN=%E4%BA%9A%E9%A9%AC%E9%80%8A%E7%BD%91%E7%AB%99&ref=nb_sb_noss")
}

func Test_getPageInfo(test *testing.T) {
	url := "/s?k=programming+books&__mk_zh_CN=%E4%BA%9A%E9%A9%AC%E9%80%8A%E7%BD%91%E7%AB%99&qid=1616999276&ref=sr_pg_1"
	pageInfo := utils.GetUrlValueByKey(website+url, "ref")
	fmt.Println("current page in-->", url, "-->", pageInfo)
}

func Test_getNextPageStr(test *testing.T) {
	fmt.Println(utils.GetNextPageStr("sr_pg_1"))
}
