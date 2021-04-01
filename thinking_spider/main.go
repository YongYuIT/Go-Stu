package main

import (
	"thinking_spider/cert_check"
	"thinking_spider/config"
	"thinking_spider/database"
	"thinking_spider/spider"
	"thinking_spider/utils"
)

func main() {

	cert_check.EnvCheck()

	priceLevelDataSpider := spider.GetPriceLevelDataSpider()
	priceLevelDataSpider.Config.MaxDeep = 2
	priceLevelDataSpider.BuildStartUrl(func(spiderConfig *config.SpiderConfig) string {
		startPage := utils.GetUrlWithKVs(spiderConfig.WebSite, []string{"k=" + utils.GetKeyWords(spiderConfig.KeyWords), "ref=nb_sb_noss"})
		priceLevelDataSpider.Ctrl.SetCookies(startPage, spiderConfig.Cookies)
		return startPage
	})

	priceLevelDataSpider.StartSpider()
	database.CloseDB()
}
