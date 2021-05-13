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
	cookie1.Value = "144-7059104-3803242"
	cookie1.Domain = ".amazon.com"
	cookie1.Path = "/"

	cookie2 := &http.Cookie{}
	cookie2.Name = "session-id-time"
	cookie2.Value = "2082787201l"
	cookie2.Domain = ".amazon.com"
	cookie2.Path = "/"

	cookie3 := &http.Cookie{}
	//cookie3.Name = "i18n-prefs"
	//cookie3.Value = "USD"
	//cookie3.Domain = ".amazon.com"
	//cookie3.Path = "/"

	cookie4 := &http.Cookie{}
	cookie4.Name = "ubid-main"
	cookie4.Value = "133-7766908-2678768"
	cookie4.Domain = ".amazon.com"
	cookie4.Path = "/"

	cookie5 := &http.Cookie{}
	cookie5.Name = "session-token"
	cookie5.Value = "byYSLRH8Dwpq21F/CV89Roc4nfIjB3tsyvHw6Xnfi6Bcsqas+t2imYhgvegujL+R4Zn+/Bdsk79/vr/hrK1vASQX0fttKNO9ANB516CDi7kYAVX2tY1GtB4HyxZuABYWyI5M85gGMwy+n6kntFq7vziNGpi7h7o0xSTwFR8OIVMDzEQStOjNirqiW9qMtKAc8sMG7UkSo+Vpad2n6waB/SR7RfRVxKpCF1/CzCjShXNPW3y//hQUIXukhYLPMZwA"
	cookie5.Domain = ".amazon.com"
	cookie5.Path = "/"

	cookie6 := &http.Cookie{}
	//cookie6.Name = "lc-main"
	//cookie6.Value = "en_US"
	//cookie6.Domain = ".amazon.com"
	//cookie6.Path = "/"

	cookie7 := &http.Cookie{}
	cookie7.Name = "csm-hit"
	cookie7.Value = "tb:s-H7QKF56YRZSWSZSRJ10M|1620894856592&t:1620894858016&adb:adblk_no"
	cookie7.Domain = "www.amazon.com"
	cookie7.Path = "/"

	cookies := []*http.Cookie{}
	if 1 == 0 {
		cookies = append(cookies, cookie1, cookie2, cookie3, cookie4, cookie5, cookie6, cookie7)
	} else {
		cookies = append(cookies, cookie1, cookie2, cookie4)
	}
	startUrl := "https://www.amazon.com/Whiteswallow-Reusable-Stainless-Whiskey-Diamond/dp/B07B3VZZ3K/ref=sr_1_12?dchild=1&keywords=Diamond+Whiskey+Chilling+Stones&qid=1620888028&refinements=p_36%3A1253523011&rnid=386465011&s=home-garden&sr=1-12&thk_p_tag=Under+%2425"
	c.SetCookies(startUrl, cookies)

	c.Visit(startUrl)
}
