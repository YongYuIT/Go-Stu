package main

import (
	"flag"
	"strings"
	"thinking_spider/cert_check"
	"thinking_spider/config"
	"thinking_spider/database"
	"thinking_spider/model"
	"thinking_spider/spider"
	"thinking_spider/utils"
)

var (
	task      string
	key_words string
)

const (
	KEYWORD_TASK = "KEYWORD_TASK"
	DETAIL_TASK  = "DETAIL_TASK"
)

func init() {
	flag.StringVar(&key_words, "k", "", "set keywords")
	flag.StringVar(&task, "t", KEYWORD_TASK, "set task")
}

func main() {

	cert_check.EnvCheck()

	flag.Parse()
	defer database.CloseDB()

	if strings.EqualFold(task, KEYWORD_TASK) {
		doKeyWorkTask()
	}

	if strings.EqualFold(task, DETAIL_TASK) {
		doDetailTask()
	}
}

func doDetailTask() {

	if !strings.EqualFold(key_words, "") {
		config.CurrentDefaultConfig.KeyWords = key_words
	}

	asinUrl := model.GetUrlByKeyWords(config.CurrentDefaultConfig.KeyWords)
	for _, asinUrl := range *asinUrl {
		detailSpider := spider.GetDetailDataSpider()
		detailSpider.Config.MaxDeep = 1
		detailSpider.BuildStartUrl(func(spiderConfig *config.SpiderConfig) string {
			startPage := config.CurrentDefaultConfig.WebSite + asinUrl.DetialUrl
			detailSpider.Ctrl.SetCookies(startPage, spiderConfig.Cookies)
			return startPage
		})
		detailSpider.StartSpider()
	}
}

func doKeyWorkTask() {

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
