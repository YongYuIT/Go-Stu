package config

import (
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"reflect"
	"strings"
)

var CurrentDefaultConfig = &SpiderConfig{}
var DBConn = &DBConfig{}

func init() {
	fmt.Println("start init config")

	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath("./")

	if err := config.ReadInConfig(); err != nil {
		fmt.Println("read config file error-->", err)
		os.Exit(1)
	}

	var database = config.Get("database").(map[string]interface{})
	autoSetValueFromConfig(DBConn, database)

	if config.Get("apply") == nil {
		return
	}
	var apply = config.Get("apply").(string)
	fmt.Println("apply-->", apply)
	applyConfig := config.Get(apply).(map[string]interface{})
	autoSetValueFromConfig(CurrentDefaultConfig, applyConfig)

	cookies := applyConfig["cookies"].(map[string]interface{})
	CurrentDefaultConfig.Cookies = autoSetCookies(cookies)

	priceLevelConfig := applyConfig["price_level"].(map[string]interface{})
	CurrentDefaultConfig.PriceLevelConfig = &PriceLevelConfig{}
	autoSetValueFromConfig(CurrentDefaultConfig.PriceLevelConfig, priceLevelConfig)

	itemsConfig := applyConfig["items"].(map[string]interface{})
	CurrentDefaultConfig.ItemsConfig = &ItemsConfig{}
	autoSetValueFromConfig(CurrentDefaultConfig.ItemsConfig, itemsConfig)

	itemConfig := itemsConfig["item"].(map[string]interface{})
	CurrentDefaultConfig.ItemsConfig.Item = &Item{}
	autoSetValueFromConfig(CurrentDefaultConfig.ItemsConfig.Item, itemConfig)
	fmt.Println("get init config-->", CurrentDefaultConfig.ItemsConfig.Item)
}

func autoSetCookies(cookiesConfig map[string]interface{}) []*http.Cookie {
	cookies := []*http.Cookie{}
	sidCookieConfig := cookiesConfig["session_id"].(map[string]interface{})
	sidCookie := &http.Cookie{}
	autoSetCookie(sidCookie, sidCookieConfig)
	cookies = append(cookies, sidCookie)

	umainCookieConfig := cookiesConfig["ubid_main"].(map[string]interface{})
	umainCookie := &http.Cookie{}
	autoSetCookie(umainCookie, umainCookieConfig)
	cookies = append(cookies, umainCookie)
	return cookies
}

func autoSetCookie(cookie *http.Cookie, cookieConfig map[string]interface{}) {
	val := reflect.ValueOf(cookie)
	val = val.Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		fieldName := field.Name
		fieldValue := val.FieldByName(fieldName)
		configVal := cookieConfig[strings.ToLower(fieldName)]
		if configVal == nil {
			continue
		}
		value := configVal.(string)
		fieldValue.Set(reflect.ValueOf(value))
	}
}

func autoSetValueFromConfig(tagObg interface{}, config map[string]interface{}) {
	val := reflect.ValueOf(tagObg)
	val = val.Elem()
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		fileType := field.Type.Name()
		tagConfig := field.Tag.Get("config")
		fieldName := field.Name
		fieldValue := val.FieldByName(fieldName)
		tagVal := config[tagConfig]
		if tagVal == nil {
			continue
		}
		if strings.EqualFold(fileType, "string") {
			fieldValue.Set(reflect.ValueOf(config[tagConfig].(string)))
		} else if strings.EqualFold(fileType, "int") {
			fieldValue.Set(reflect.ValueOf(config[tagConfig].(int)))
		}
	}
}
