package simple_test

import (
	"fmt"
	"github.com/gocolly/colly"
	"os"
	"strings"
	"testing"
	"thinking_spider/utils"
	"time"
)

const website = "https://www.amazon.com"

func Test_when_robot(test *testing.T) {
	for i := 0; i < 500; i++ {
		go go_to_next_page("programing+book")
	}

	for true {
		fmt.Println("waiting")
		time.Sleep(time.Duration(2) * time.Second)
	}
}

func Test_go_to_next_page(test *testing.T) {
	go_to_next_page("programing+book")
}

func go_to_next_page(key_word string) {
	c := colly.NewCollector(
		colly.MaxDepth(5),
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
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
		fmt.Fprintf(file, string(response.Body))
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

	c.OnHTML("div[class='a-container a-padding-double-large'] p[class='a-last']", func(element *colly.HTMLElement) {
		if strings.Contains(element.Text, "robot") {
			fmt.Println("ant robot when req-->",
				element.Request.URL.String(),
				"\n time:", time.Now().Format(time.RFC850),
				"\n keyword: ", key_word)
		}
	})

	c.OnError(func(response *colly.Response, err error) {
		fmt.Println("find a error when req-->",
			response.Request.URL.String(),
			"\n time:", time.Now().Format(time.RFC850),
			"\n keyword: ", key_word,
			"\n details for error", err)
	})

	c.Visit(website + "/s?k=" + key_word)
}

func Test_getPageInfo(test *testing.T) {
	url := "/s?k=programming+books&__mk_zh_CN=%E4%BA%9A%E9%A9%AC%E9%80%8A%E7%BD%91%E7%AB%99&qid=1616999276&ref=sr_pg_1"
	pageInfo := utils.GetUrlValueByKey(website+url, "ref")
	fmt.Println("current page in-->", url, "-->", pageInfo)
}

func Test_getNextPageStr(test *testing.T) {
	fmt.Println(utils.GetNextPageStr("sr_pg_1"))
}
