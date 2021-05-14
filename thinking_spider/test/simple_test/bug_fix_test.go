package simple

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"net/http"
	"os"
	"testing"
)

func Test_hello_fix(test *testing.T) {
	c := colly.NewCollector(
		colly.MaxDepth(10),
		func(collector *colly.Collector) {
			extensions.RandomUserAgent(collector)
			extensions.Referer(collector)
		})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Println("resp-code-->", response.Request.URL, "-->", response.StatusCode)
		file, err := os.Create("bug_fix_bad.html")
		if err != nil {
			return
		}
		defer file.Close()
		//fmt.Fprintf(file, string(response.Body))
		fmt.Fprintf(file, "<!-- "+response.Request.URL.String()+" -->\n")
		file.Write(response.Body)
	})

	cookie1 := &http.Cookie{}
	cookie1.Name = "session-id"
	cookie1.Value = "130-6556824-3391359"
	cookie1.Domain = ".amazon.com"
	cookie1.Path = "/"

	cookie2 := &http.Cookie{}
	cookie2.Name = "session-id-time"
	cookie2.Value = "2082787201l"
	cookie2.Domain = ".amazon.com"
	cookie2.Path = "/"

	cookie3 := &http.Cookie{}
	cookie3.Name = "i18n-prefs"
	cookie3.Value = "USD"
	cookie3.Domain = ".amazon.com"
	cookie3.Path = "/"

	cookie4 := &http.Cookie{}
	cookie4.Name = "ubid-main"
	cookie4.Value = "134-9616237-3214505"
	cookie4.Domain = ".amazon.com"
	cookie4.Path = "/"

	cookie5 := &http.Cookie{}
	cookie5.Name = "session-token"
	cookie5.Value = "NsqHDpg+qsbxDpQywkh+DHJjyAteAtcugaQQsZjHjE7a33HeMK7I1oKpMI+MN5bUYXsSkfbMkzS5D0yXJMsgfIcO0OeVlQDQxpYRiUF8zCzYhj6/qVwrJ2jdAoqNcWdJdFjbmrlXfShcGOXJMIMOIAJocdsg1azdbvVvh0Ry/uSfR0DRMiTr2u4/kQKyA9U6"
	cookie5.Domain = ".amazon.com"
	cookie5.Path = "/"

	cookie6 := &http.Cookie{}
	cookie6.Name = "lc-main"
	cookie6.Value = "en_US"
	cookie6.Domain = ".amazon.com"
	cookie6.Path = "/"

	cookie7 := &http.Cookie{}
	cookie7.Name = "csm-hit"
	cookie7.Value = "tb:s-G6GZPZWZAQ4CXWS47YA9|1620958707574&t:1620958708154&adb:adblk_no"
	cookie7.Domain = "www.amazon.com"
	cookie7.Path = "/"

	cookies := []*http.Cookie{}
	if 1 == 0 {
		cookies = append(cookies, cookie1, cookie2, cookie3, cookie4, cookie5, cookie6, cookie7)
	} else {
		cookies = append(cookies, cookie1, cookie2, cookie3, cookie4, cookie5, cookie6, cookie7)
	}
	startUrl := "https://www.amazon.com/s?k=wine+corks&ref=nb_sb_noss"
	c.SetCookies(startUrl, cookies)

	c.Visit(startUrl)
}
