package handler

import (
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
	"thinking_spider/config"
	"thinking_spider/database"
	"thinking_spider/model"
	"thinking_spider/utils"
)

func GetProductItemHandler(config *config.SpiderConfig) colly.HTMLCallback {
	return func(element *colly.HTMLElement) {
		element.ForEach(config.ItemsConfig.ProductItemQue, func(i int, eItemDiv *colly.HTMLElement) {
			record := model.NewKeyWordProdRecord()
			record.KeyWord = config.KeyWords
			record.PriceLevel = config.CurrentPriceLevel
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
			record.Desc = eItemDiv.ChildAttr(config.ItemsConfig.Item.ItemDescQue, config.ItemsConfig.Item.ItemDescAttr)
			record.Price = utils.GetPrice(eItemDiv.ChildText(config.ItemsConfig.Item.ItemPriceQue))
			record.Sales, err = strconv.Atoi(strings.ReplaceAll(eItemDiv.ChildText(config.ItemsConfig.Item.ItemSalesQue), ",", ""))
			if err != nil {
				fmt.Println("get item error-->", err)
				record.Sales = -1
			}
			saveRecord(record)
		})
	}
}

func saveRecord(record *model.KeyWordProdRecord) {
	fmt.Println("record-->", record)
	database.SaveKeyWordProdRecord(record)
}