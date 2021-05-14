package main

import (
	"flag"
	"fmt"
	"strings"
	"thinking_spider/cert_check"
	"thinking_spider/config"
	"thinking_spider/database"
	"thinking_spider/model"
	"thinking_spider/spider"
	"thinking_spider/utils"
	"time"
)

var (
	task      string
	key_words string
)

const (
	KEYWORD_TASK     = "KEYWORD_TASK"
	DETAIL_TASK      = "DETAIL_TASK"
	NEW_RELEASE_TASK = "NEW_RELEASE_TASK"
	PATENTS_TASK     = "PATENTS_TASK"
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
	} else if strings.EqualFold(task, DETAIL_TASK) {
		doDetailTask()
	} else if strings.EqualFold(task, NEW_RELEASE_TASK) {
		doReleaseTask()
	} else if strings.EqualFold(task, PATENTS_TASK) {
		doPatentsTask()
	} else {
		fmt.Println("not task matches")
	}
}

func doPatentsTask() {

	if !strings.EqualFold(key_words, "") {
		config.CurrentDefaultConfig.KeyWords = key_words
	}

	mainspider := spider.GetPatentSpider()
	mainspider.BuildStartUrl(func(spiderConfig *config.SpiderConfig) string {
		return mainspider.Config.KeyWords
	})
	mainspider.StartSpider()
}

func doReleaseTask() {

	config.CurrentDefaultConfig.TaskIndex = (int)(time.Now().Unix())

	homeAndKitchen := make(map[string]interface{})
	config.InitHomeAndKitchen(homeAndKitchen)

	newReleaseTypesSpider := spider.GetNewReleaseTypeSpider()
	newReleaseTypesSpider.BuildStartUrl(func(spiderConfig *config.SpiderConfig) string {
		startUrl := newReleaseTypesSpider.Config.NewRelease
		newReleaseTypesSpider.SetPageValue(startUrl, "tree", config.GetMapInMap(spiderConfig.KeyWords, homeAndKitchen))
		cookies := spiderConfig.GetCookie()
		fmt.Println("set cookies-->", cookies[0].Name, "-->", cookies[0].Value)
		newReleaseTypesSpider.Ctrl.SetCookies(startUrl, cookies)
		return startUrl
	})
	newReleaseTypesSpider.StartSpider()
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
			cookies := spiderConfig.GetCookie()
			fmt.Println("set cookies-->", cookies[0].Name, "-->", cookies[0].Value)
			detailSpider.SetPageValue(startPage, "asin", asinUrl.Asin)
			detailSpider.Ctrl.SetCookies(startPage, cookies)
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
		cookies := spiderConfig.GetCookie()
		fmt.Println("set cookies-->", cookies[0].Name, "-->", cookies[0].Value)
		priceLevelDataSpider.Ctrl.SetCookies(startPage, cookies)
		return startPage
	})
	priceLevelDataSpider.StartSpider()
}
