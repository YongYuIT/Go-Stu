package handler

import (
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
	"thinking_spider/model"
	"thinking_spider/spider_interface"
	"thinking_spider/utils"
)

func GetProductItemHandler(spider *spider_interface.Spider) colly.HTMLCallback {
	return func(element *colly.HTMLElement) {
		config := spider.Config
		element.ForEach(config.ItemsConfig.ProductItemQue, func(i int, eItemDiv *colly.HTMLElement) {
			record := model.NewKeyWordProdRecord()
			record.KeyWord = config.KeyWords
			record.PriceLevel = config.CurrentPriceLevel
			record.DeliverTo = spider.GetPageValue(element.Request.URL.String(), spider_interface.REGION_NAME).(string)
			pageinfo := utils.GetUrlValueByKey(eItemDiv.Request.URL.String(), config.PagesKey)
			if strings.EqualFold(pageinfo, config.PageStartTag) || strings.Contains(pageinfo, "sr_nr_") {
				record.Page = 1
			} else {
				record.Page = utils.GetPageNum(utils.GetUrlValueByKey(eItemDiv.Request.URL.String(), config.PagesKey))
			}
			record.Asin = eItemDiv.Attr(config.ItemsConfig.Item.ItemAsinAttr)
			pageIndex, err := strconv.Atoi(eItemDiv.Attr(config.ItemsConfig.Item.ItemIndexAttr))
			if err != nil {
				fmt.Println("get item error-->", err)
			}
			record.PageIndex = pageIndex
			record.Uuid = eItemDiv.Attr(config.ItemsConfig.Item.ItemUUIDAttr)
			record.Titles = eItemDiv.ChildAttr(config.ItemsConfig.Item.ItemDescQue, config.ItemsConfig.Item.ItemDescAttr)
			record.Price = utils.GetPrice(eItemDiv.ChildText(config.ItemsConfig.Item.ItemPriceQue))
			record.Ratings, err = strconv.Atoi(strings.ReplaceAll(eItemDiv.ChildText(config.ItemsConfig.Item.ItemRatingsQue), ",", ""))
			record.Starts = utils.GetStarts(eItemDiv.ChildAttr(config.ItemsConfig.Item.ItemStartsQue, "aria-label"))
			record.DetialUrl = eItemDiv.ChildAttr(config.ItemsConfig.Item.ItemDetailUrlQue, "href")
			record.MainPicUrl = eItemDiv.ChildAttr(config.ItemsConfig.Item.ItemImgUrlQue, "src")

			if err != nil {
				fmt.Println("get item error-->", err)
				record.Ratings = -1
			}
			saveRecord(record)
		})
	}
}

func saveRecord(record *model.KeyWordProdRecord) {
	fmt.Println("record-->", record)
	model.SaveKeyWordProdRecord(record)
}
