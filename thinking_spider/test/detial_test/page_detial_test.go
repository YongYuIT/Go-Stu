package detial_test

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"
)

func Test_hello_fix(test *testing.T) {
	c := colly.NewCollector(
		colly.MaxDepth(1),
		func(collector *colly.Collector) {
			extensions.RandomUserAgent(collector)
			extensions.Referer(collector)
		},
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
		//fmt.Fprintf(file, string(response.Body))
		file.Write(response.Body)
	})

	c.OnHTML("ul[class='a-unordered-list a-vertical a-spacing-mini'] li[id!='replacementPartsFitmentBullet'] span[class='a-list-item']", func(element *colly.HTMLElement) {
		desc := strings.TrimSpace(element.Text)
		fmt.Println("-----------------------start")
		fmt.Println("desc-->", desc)
		fmt.Println("-----------------------end")
	})
	c.OnHTML("input[type='hidden'][id='ftSelectAsin']", func(element *colly.HTMLElement) {
		asin := element.Attr("value")
		fmt.Println("asin-->", asin)
	})
	c.OnHTML("span[class='a-size-small'] a#sellerProfileTriggerId", func(element *colly.HTMLElement) {
		sold := element.Text
		fmt.Println("sold by-->", sold)
	})
	c.OnHTML("input[type='hidden'][id='ftSelectMerchant']", func(element *colly.HTMLElement) {
		soldid := element.Attr("value")
		fmt.Println("soldid-->", soldid)
	})
	c.OnHTML("div#productDescription[class='a-section a-spacing-small'] p", func(element *colly.HTMLElement) {
		element.DOM.Contents().Each(func(i int, selection *goquery.Selection) {
			if selection.Is("br") {
				fmt.Println("pdesc-->", selection.Nodes[0].PrevSibling.Data)
				fmt.Println("pdesc-->", selection.Nodes[0].NextSibling.Data)
			}
		})
	})

	cookie1 := &http.Cookie{}
	cookie1.Name = "session-id"
	cookie1.Value = "147-6489143-1689104"
	cookie1.Domain = ".amazon.com"
	cookie1.Path = "/"

	cookie4 := &http.Cookie{}
	cookie4.Name = "ubid-main"
	cookie4.Value = "132-6250557-3411723"
	cookie4.Domain = ".amazon.com"
	cookie4.Path = "/"

	cookies := []*http.Cookie{}
	cookies = append(cookies, cookie1, cookie4)

	//startUrl := "https://www.amazon.com/QWERBAM-Multifunction-Double-Sided-Removable-Traceless/dp/B08H57B5K4/ref=sr_1_135?dchild=1&keywords=traceless+double-sided+tape&qid=1618384945&refinements=p_36%3A1253551011&rnid=386479011&s=office-products&sr=1-135&thk_p_tag=%2425+to+%2450"
	startUrl := "https://www.amazon.com/MDYH-Stainless-Barbecue-Roasting-Shovels/dp/B08KQD15NS/ref=sr_1_13?dchild=1&keywords=frying+shovel+clip+fried+fish+shovel+steak+shovel+home+kitchen+spatula+fried+steak+clip+pancake&qid=1618366915&refinements=p_36%3A2661612011&rnid=2661611011&s=lawn-garden&sr=1-13&thk_p_tag=Under+%2425"

	c.SetCookies(startUrl, cookies)

	c.Visit(startUrl)
}
