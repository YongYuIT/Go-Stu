package test

import (
	"testing"
	"thinking_spider/config"
	"thinking_spider/database"
	"thinking_spider/spider"
	"thinking_spider/utils"
)

func Test_getDataSpider(test *testing.T) {
	priceLevelDataSpider := spider.GetPriceLevelDataSpider()
	priceLevelDataSpider.Config.MaxDeep = 2
	priceLevelDataSpider.BuildStartUrl(func(spiderConfig *config.SpiderConfig) string {
		return utils.GetUrlWithKVs(spiderConfig.WebSite, []string{"k=" + utils.GetKeyWords(spiderConfig.KeyWords), "ref=nb_sb_noss"})

	})
	priceLevelDataSpider.StartSpider()
	database.CloseDB()
}
