package handler

import (
	"fmt"
	"strings"
	"thinking_spider/database"
	"thinking_spider/utils"

	"github.com/gocolly/colly"
	"strconv"
	"thinking_spider/config"
	"thinking_spider/model"
)

func ProductItemHandler(element *colly.HTMLElement) {
	element.ForEach(config.CurrentSprierConfig.Items.ProductItemQue, func(i int, eItemDiv *colly.HTMLElement) {
		record := model.NewKeyWordProdRecord()
		record.KeyWord = config.CurrentSprierConfig.KeyWords
		pageinfo := utils.GetPageInfo(eItemDiv.Request.URL.String(), config.CurrentSprierConfig.PagesKey)
		if strings.EqualFold(pageinfo, config.CurrentSprierConfig.PageStartTag) {
			record.Page = 1
		} else {
			record.Page = utils.GetPageNum(utils.GetPageInfo(eItemDiv.Request.URL.String(), config.CurrentSprierConfig.PagesKey))
		}
		record.Asin = eItemDiv.Attr(config.CurrentSprierConfig.Items.Item.ItemAsinAttr)
		pageIndex, err := strconv.Atoi(eItemDiv.Attr(config.CurrentSprierConfig.Items.Item.ItemIndex))
		if err != nil {
			return
		}
		record.PageIndex = pageIndex
		record.Uuid = eItemDiv.Attr(config.CurrentSprierConfig.Items.Item.ItemUUIDAttr)
		record.Desc = eItemDiv.ChildAttr(config.CurrentSprierConfig.Items.Item.ItemDescQue, config.CurrentSprierConfig.Items.Item.ItemDescAttr)
		record.Sales, err = strconv.Atoi(eItemDiv.ChildText(config.CurrentSprierConfig.Items.Item.ItemSalesQue))
		if err != nil {
			return
		}
		saveRecord(record)
	})
}

func saveRecord(record *model.KeyWordProdRecord) {
	fmt.Println("record-->", record)
	database.SaveKeyWordProdRecord(record)
}
