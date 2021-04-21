package handler

import (
	"github.com/gocolly/colly"
	"strconv"
	"strings"
	"thinking_spider/model"
	"thinking_spider/spider_interface"
)

func GetNewReaseItemHandler(spider *spider_interface.Spider) colly.HTMLCallback {
	return func(element *colly.HTMLElement) {
		element.ForEach("li[class='zg-item-immersion']", func(i int, listEle *colly.HTMLElement) {
			release := model.NewNewReleaseProdRecord()
			types := strings.Split(spider.Config.KeyWords, "##")
			if len(types) > 0 {
				release.Type1 = types[0]
			}
			if len(types) > 1 {
				release.Type2 = types[1]
			}
			if len(types) > 2 {
				release.Type3 = types[2]
			}
			if len(types) > 3 {
				release.Type4 = types[3]
			}
			indexStr := listEle.ChildText("span[class='zg-badge-text']")[1:]
			index, err := strconv.Atoi(indexStr)
			if err == nil {
				release.Index = index
			}
			asin := listEle.ChildAttr("span[class='aok-inline-block zg-item'] a[href]", "href")
			release.Asin = asin[strings.Index(asin, "dp/")+3 : strings.Index(asin, "dp/")+13]
			release.Titles = listEle.ChildText("div[aria-hidden='true']")
			ratings, err := strconv.Atoi(listEle.ChildText("a[class='a-size-small a-link-normal']"))
			if err == nil {
				release.Ratings = ratings
			}
			startStr := listEle.ChildText("span[class='a-icon-alt']")
			if strings.Contains(startStr, "out") {
				startStr = strings.TrimSpace(startStr[:strings.Index(startStr, "out")])
				start, err := strconv.ParseFloat(startStr, 32)
				if err == nil {
					release.Starts = float32(start)
				}
			}
			listEle.ForEach("span[class='p13n-sc-price']", func(i int, element *colly.HTMLElement) {
				if strings.EqualFold(release.Price, "") {
					release.Price = element.Text
				} else {
					release.Price += "~" + element.Text
				}
			})
			devTo := spider.GetPageValue(element.Request.URL.String(), spider_interface.REGION_NAME)
			if devTo != nil {
				release.DeliverTo = devTo.(string)
			}
			release.DetialUrl = asin
			release.MainPicUrl = listEle.ChildAttr("img", "src")
			model.SaveNewReleaseProdRecord(release)
		})
	}
}
