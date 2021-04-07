package utils

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"thinking_spider/config"
)

func GetUrlWithKVs(host string, kvs []string) string {
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

func GetUrlValueByKey(urlstr string, key string) string {
	thisUrl, err := url.Parse(urlstr)
	if err != nil {
		return ""
	}
	kvs, err := url.ParseQuery(thisUrl.RawQuery)
	if err != nil {
		return ""
	}
	if len(kvs[key]) > 0 {
		return kvs[key][0]
	} else {
		return ""
	}
}

func GetPageNum(str string) int {
	page := str[len(config.CurrentDefaultConfig.PageUrlTag):]
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

func GetPrice(str string) float32 {
	if strings.Contains(str, "¥") {
		str = str[len("¥"):]
	} else if strings.Contains(str, "$") {
		str = str[len("$"):]
	}
	value, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return -1
	}
	return float32(value)
}

func GetStarts(str string) float32 {
	if !strings.Contains(str, "out") {
		return -1
	}
	str = str[:strings.Index(str, "out")]
	str = strings.TrimSpace(str)
	value, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return -1
	}
	return float32(value)
}
