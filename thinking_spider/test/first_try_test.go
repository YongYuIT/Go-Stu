package test

import (
	"fmt"
	"github.com/gocolly/colly"
	"testing"
)

func Test_hello(test *testing.T) {
	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("http://go-colly.org/")
}

func Test_get_product_list(test *testing.T) {
	c := colly.NewCollector()
	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		fmt.Println("-----------------start")
		fmt.Println(e.Text)
		fmt.Println("-----------------end")
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.Visit("https://www.amazon.com/gp/site-directory/")
}

func Test_get_product_list1(test *testing.T) {
	c := colly.NewCollector()
	// Find and visit all links
	c.OnHTML("div#nav-xshop", func(e *colly.HTMLElement) {
		e.ForEach("a[href]", func(i int, element *colly.HTMLElement) {
			fmt.Println(i, element.Request.URL, element.Text)
		})
	})
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	c.Visit("https://www.amazon.com/gp/site-directory/")
}

func Test_get_product_key_word(test *testing.T) {
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Println("resp-code-->", response.StatusCode)
	})

	c.OnHTML("div[class='s-main-slot s-result-list s-search-results sg-row']", func(element *colly.HTMLElement) {
		fmt.Println("current page is-->", element.Request.URL.Path)
		element.ForEach("div[data-asin]", func(i int, eItemDiv *colly.HTMLElement) {
			fmt.Println(eItemDiv.Attr("data-asin"), "-->", eItemDiv.Attr("data-uuid"))
			fmt.Println("prod desc-->", eItemDiv.ChildAttr("img[class='s-image']", "alt"))
			fmt.Println("sales volume-->", eItemDiv.ChildText("a[class='a-link-normal'] span[class='a-size-base']"))
		})
	})

	c.OnHTML("div[class='a-text-center']", func(element *colly.HTMLElement) {
		fmt.Println("-----------------------------------page start")
		element.ForEach("li[class=a-normal] a[href]", func(i int, element *colly.HTMLElement) {
			fmt.Println("get pages-->", element.Attr("href"))
		})
		fmt.Println("-----------------------------------")
		element.ForEach("li[class=a-selected] a[href]", func(i int, element *colly.HTMLElement) {
			fmt.Println("get pages-->", element.Attr("href"))
		})
		fmt.Println("-----------------------------------page end")
	})

	//c.Visit("https://www.amazon.com/s/ref=nb_sb_noss?url=search-alias%3Daps&field-keywords=programming+books")
	c.Visit("https://www.amazon.cn/s?k=programming+books&__mk_zh_CN=%E4%BA%9A%E9%A9%AC%E9%80%8A%E7%BD%91%E7%AB%99&ref=nb_sb_noss")
}
