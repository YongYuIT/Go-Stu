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
	cookie1.Value = "143-6303151-7535421"
	cookie1.Domain = ".amazon.com"
	cookie1.Path = "/"

	cookie2 := &http.Cookie{}
	//cookie2.Name = "session-id-time"
	//cookie2.Value = "2082787201l"
	//cookie2.Domain = ".amazon.com"
	//cookie2.Path = "/"

	cookie3 := &http.Cookie{}
	//cookie3.Name = "i18n-prefs"
	//cookie3.Value = "USD"
	//cookie3.Domain = ".amazon.com"
	//cookie3.Path = "/"

	cookie4 := &http.Cookie{}
	cookie4.Name = "ubid-main"
	cookie4.Value = "134-4873337-3736159"
	cookie4.Domain = ".amazon.com"
	cookie4.Path = "/"

	cookie5 := &http.Cookie{}
	//cookie5.Name = "session-token"
	//cookie5.Value = "Om3pjbDhhLi7HHPplHkj1KbFktFSNn4ZMAW0fpJw6P/wZgyMWOdtnoSU5R6lPvwJwe+1F0cD3FshN3BeiE1cMJ2afmqXomU6yf0Nag3jmrFxNsb1r5BFSoBiuGH8uJXmYIWy3qivA+RrewCcRpwcZpB50xCrKlMgooyVfBEnonc3acqLTG3tFh3CxNPiCwn5pUPbaCwFg5E/q2IlYxDd4P0avHLK6nMZDfJUd8bav88ZKnxm7gUCBJYWXLbw/igW"
	//cookie5.Domain = ".amazon.com"
	//cookie5.Path = "/"

	cookie6 := &http.Cookie{}
	//cookie6.Name = "lc-main"
	//cookie6.Value = "en_US"
	//cookie6.Domain = ".amazon.com"
	//cookie6.Path = "/"

	cookie7 := &http.Cookie{}
	//cookie7.Name = "csm-hit"
	//cookie7.Value = "tb:s-CF8R256R2EBE27MHFQ6J|1617176843169&t:1617176843169&adb:adblk_no"
	//cookie7.Domain = "www.amazon.com"
	//cookie7.Path = "/"

	cookies := []*http.Cookie{}
	cookies = append(cookies, cookie1, cookie2, cookie3, cookie4, cookie5, cookie6, cookie7)

	startUrl := "https://www.amazon.com/s?k=cheese+spreaders&ref=nb_sb_noss"
	c.SetCookies(startUrl, cookies)

	c.Visit(startUrl)
}
