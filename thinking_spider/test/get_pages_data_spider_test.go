package test

import (
	"testing"
	"thinking_spider/config"
	"thinking_spider/spider"
	"thinking_spider/utils"
)

func Test_use_spider_get_pages_data(test *testing.T) {
	var thisConfig = *config.CurrentDefaultConfig
	thisConfig.MaxDeep = 4
	thisConfig.DelaySpider = 0
	spider := spider.GetPagesDataSpider()
	spider.Config.MaxDeep = 4
	spider.Config.DelaySpider = 0
	spider.BuildStartUrl(func(spiderConfig *config.SpiderConfig) string {
		return utils.GetUrlWithKVs(thisConfig.WebSite, []string{"k=" + utils.GetKeyWords(thisConfig.KeyWords), "ref=nb_sb_noss"})
	})
	spider.StartSpider()
}
