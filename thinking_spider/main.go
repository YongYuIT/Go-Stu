package main

import (
	"flag"
	"strings"
	"thinking_spider/cert_check"
	"thinking_spider/config"
	"thinking_spider/database"
	"thinking_spider/spider"
	"thinking_spider/utils"
)

var (
	key_words string
)

func init() {
	flag.StringVar(&key_words, "k", "", "set keywords")
}

func main() {

	cert_check.EnvCheck()

	flag.Parse()
	defer database.CloseDB()

	if !strings.EqualFold(key_words, "") {
		config.CurrentDefaultConfig.KeyWords = key_words
	}

	priceLevelDataSpider := spider.GetPriceLevelDataSpider()
	priceLevelDataSpider.Config.MaxDeep = 2
	priceLevelDataSpider.BuildStartUrl(func(spiderConfig *config.SpiderConfig) string {
		startPage := utils.GetUrlWithKVs(spiderConfig.WebSite, []string{"k=" + utils.GetKeyWords(spiderConfig.KeyWords), "ref=nb_sb_noss"})
		priceLevelDataSpider.Ctrl.SetCookies(startPage, spiderConfig.Cookies)
		return startPage
	})

	priceLevelDataSpider.StartSpider()
}
