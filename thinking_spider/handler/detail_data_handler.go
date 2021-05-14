package handler

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"strings"
	"thinking_spider/model"
	"thinking_spider/spider_interface"
)

func doSaveRecord(spider *spider_interface.Spider, record *model.ProdDetailRecord, element *colly.HTMLElement) {
	if strings.EqualFold(record.Asin, "") {
		if spider.GetPageValue(element.Request.URL.String(), "asin") != nil {
			record.Asin = spider.GetPageValue(element.Request.URL.String(), "asin").(string)
		}
	}
	model.SaveProdDetailRecord(record)
}

func Get5DescsHandler(spider *spider_interface.Spider, record *model.ProdDetailRecord) (string, colly.HTMLCallback) {
	callback := func(element *colly.HTMLElement) {
		if record.Desc1 == "" {
			record.Desc1 = strings.TrimSpace(element.Text)
		} else if record.Desc2 == "" {
			record.Desc2 = strings.TrimSpace(element.Text)
		} else if record.Desc3 == "" {
			record.Desc3 = strings.TrimSpace(element.Text)
		} else if record.Desc4 == "" {
			record.Desc4 = strings.TrimSpace(element.Text)
		} else if record.Desc5 == "" {
			record.Desc5 = strings.TrimSpace(element.Text)
		}
		doSaveRecord(spider, record, element)
	}
	return spider.Config.DetailsConfig.DescsQue, callback
}

func GetAsin(spider *spider_interface.Spider, record *model.ProdDetailRecord) (string, colly.HTMLCallback) {
	callback := func(element *colly.HTMLElement) {
		record.Asin = element.Attr("value")
		if strings.EqualFold(record.Asin, "") {
			record.Asin = spider.GetPageValue(element.Request.URL.String(), "asin").(string)
		}
		doSaveRecord(spider, record, element)
	}
	return spider.Config.DetailsConfig.AsinQue, callback
}

func GetSoldBy(spider *spider_interface.Spider, record *model.ProdDetailRecord) (string, colly.HTMLCallback) {
	callback := func(element *colly.HTMLElement) {
		value := element.ChildText("div[class='a-row'] span[class='a-size-small']")
		if strings.Contains(element.ChildText("span[class='a-size-small a-color-tertiary']"), "Ships") {
			record.ShipsFrom = value
		} else if strings.Contains(element.ChildText("span[class='a-size-small a-color-tertiary']"), "Sold") {
			record.SoldBy = value
		}
		doSaveRecord(spider, record, element)
	}
	return spider.Config.DetailsConfig.SoldByQue, callback
}

func GetProdDesc(spider *spider_interface.Spider, record *model.ProdDetailRecord) (string, colly.HTMLCallback) {

	callback := func(element *colly.HTMLElement) {
		descs := ""
		element.DOM.Contents().Each(func(i int, selection *goquery.Selection) {
			if selection.Is("br") {
				desc := ""
				if selection.Nodes[0].PrevSibling != nil {
					desc = selection.Nodes[0].PrevSibling.Data
					if !strings.Contains(descs, desc) {
						descs += desc + "\n"
					}
				}
				if selection.Nodes[0].NextSibling != nil {
					desc = selection.Nodes[0].NextSibling.Data
					if !strings.Contains(descs, desc) {
						descs += desc + "\n"
					}
				}
			}
		})
		record.ProdDesc = descs

		if strings.EqualFold(record.ProdDesc, "") {
			record.ProdDesc = strings.TrimSpace(element.Text)
		}
		doSaveRecord(spider, record, element)
	}
	return spider.Config.DetailsConfig.ProdDescQue, callback
}

func GetSoldID(spider *spider_interface.Spider, record *model.ProdDetailRecord) (string, colly.HTMLCallback) {
	callback := func(element *colly.HTMLElement) {
		record.SoldByAsin = element.Attr("value")
		doSaveRecord(spider, record, element)
	}
	return spider.Config.DetailsConfig.SoldIdQue, callback
}

func GetDataFirstAvailable(spider *spider_interface.Spider, record *model.ProdDetailRecord) (string, colly.HTMLCallback) {
	callback := func(element *colly.HTMLElement) {
		date := ""
		element.ForEach("tr", func(i int, element *colly.HTMLElement) {
			itemTitle := element.ChildText("th']")
			strings.Contains(itemTitle, "Available")
			date = strings.TrimSpace(element.ChildText("td"))
		})
		if len(date) < 40 && strings.Index(date, ",") > -1 {
			record.DateFirstAvailable = date
		}
		doSaveRecord(spider, record, element)
	}
	return "table#productDetails_detailBullets_sections1", callback
}
