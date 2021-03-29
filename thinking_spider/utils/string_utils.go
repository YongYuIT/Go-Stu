package utils

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"thinking_spider/config"
)

func GetUrl(host string, kvs []string) string {
	url := host + "/s?"
	for i, kv := range kvs {
		if i > 0 {
			url += "&"
		}
		url += kv
	}
	return url
}

func GetKeyWords(keyStr string) string {
	words := strings.Split(keyStr, " ")
	keyWords := ""
	for i, word := range words {
		keyWords += word
		if i != len(words)-1 {
			keyWords += "+"
		}
	}
	return keyWords
}

func GetPageInfo(urlstr string, key string) string {
	thisUrl, err := url.Parse(urlstr)
	if err != nil {
		return ""
	}
	kvs, err := url.ParseQuery(thisUrl.RawQuery)
	if err != nil {
		return ""
	}
	return kvs[key][0]
}

func GetPageNum(str string) int {
	page := str[len(config.CurrentSprierConfig.PageUrlTag):]
	num, err := strconv.Atoi(page)
	if err != nil {
		return -1
	}
	return num
}

func GetNextPageStr(current string) string {
	num := GetPageNum(current)
	if num == -1 {
		return ""
	}
	return fmt.Sprintf("sr_pg_%d", num+1)
}
