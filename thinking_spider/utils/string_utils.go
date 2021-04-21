package utils

import (
	"net/url"
	"strconv"
	"strings"
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
	}
	//find an other way
	if strings.Contains(urlstr, key) {
		subStr := urlstr[strings.Index(urlstr, key):]
		subStr = subStr[strings.Index(subStr, "=")+1 : strings.Index(subStr, "?")]
		return subStr
	}
	return ""
}

func GetPageNum(str string) int {
	i := len(str) - 1
	for {
		if i < 0 {
			break
		}
		if str[i] < '0' || str[i] > '9' {
			break
		}
		i--
	}
	if i < 0 {
		return -1
	} else {
		numStr := str[i+1:]
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return -1
		}
		return num
	}
}

func GetNextPageStr(current string) string {
	num := GetPageNum(current)
	numStr := strconv.Itoa(num)
	if num == -1 {
		return ""
	}
	return current[:len(current)-len(numStr)] + strconv.Itoa(num+1)
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
