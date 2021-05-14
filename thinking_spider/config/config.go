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

	p_cookies := applyConfig["cookies"]
	if p_cookies != nil {
		cookies := p_cookies.([]interface{})
		CurrentDefaultConfig.Cookies = autoSetCookies(cookies)
	}

	p_detailsConfig := applyConfig["details"]
	if p_detailsConfig != nil {
		detailsConfig := p_detailsConfig.(map[string]interface{})
		CurrentDefaultConfig.DetailsConfig = &DetailsConfig{}
		autoSetValueFromConfig(CurrentDefaultConfig.DetailsConfig, detailsConfig)
	}

	p_priceLevelConfig := applyConfig["price_level"]
	if p_priceLevelConfig != nil {
		priceLevelConfig := p_priceLevelConfig.(map[string]interface{})
		CurrentDefaultConfig.PriceLevelConfig = &PriceLevelConfig{}
		autoSetValueFromConfig(CurrentDefaultConfig.PriceLevelConfig, priceLevelConfig)
	}

	p_itemsConfig := applyConfig["items"]
	var itemsConfig map[string]interface{} = nil
	if p_itemsConfig != nil {
		itemsConfig = p_itemsConfig.(map[string]interface{})
		CurrentDefaultConfig.ItemsConfig = &ItemsConfig{}
		autoSetValueFromConfig(CurrentDefaultConfig.ItemsConfig, itemsConfig)
	}

	if itemsConfig != nil {
		itemConfig := itemsConfig["item"].(map[string]interface{})
		CurrentDefaultConfig.ItemsConfig.Item = &Item{}
		autoSetValueFromConfig(CurrentDefaultConfig.ItemsConfig.Item, itemConfig)
	}
}

func autoSetCookies(cookiesConfig []interface{} /*map[string]interface{}*/) [][]*http.Cookie {

	cookies := [][]*http.Cookie{}
	for i := range cookiesConfig {
		cookieConfig := cookiesConfig[i].(map[interface{}]interface{})["cookie"].(map[interface{}]interface{})
		cookie := []*http.Cookie{}
		sidCookieConfig := cookieConfig["session_id"].(map[interface{}]interface{})
		sidCookie := &http.Cookie{}
		autoSetCookie(sidCookie, sidCookieConfig)
		cookie = append(cookie, sidCookie)

		umainCookieConfig := cookieConfig["ubid_main"].(map[interface{}]interface{})
		umainCookie := &http.Cookie{}
		autoSetCookie(umainCookie, umainCookieConfig)
		cookie = append(cookie, umainCookie)

		sidTimeCookieConfig := cookieConfig["session_id_time"].(map[interface{}]interface{})
		sidTimeCookie := &http.Cookie{}
		autoSetCookie(sidTimeCookie, sidTimeCookieConfig)
		cookie = append(cookie, sidTimeCookie)

		i18nPrefsCookieConfig := cookieConfig["i18n_prefs"].(map[interface{}]interface{})
		i18nPrefsCookie := &http.Cookie{}
		autoSetCookie(i18nPrefsCookie, i18nPrefsCookieConfig)
		cookie = append(cookie, i18nPrefsCookie)

		sTokenCookieConfig := cookieConfig["session_token"].(map[interface{}]interface{})
		sTokenCookie := &http.Cookie{}
		autoSetCookie(sTokenCookie, sTokenCookieConfig)
		cookie = append(cookie, sTokenCookie)

		lcMainCookieConfig := cookieConfig["lc_main"].(map[interface{}]interface{})
		lcMainCookie := &http.Cookie{}
		autoSetCookie(lcMainCookie, lcMainCookieConfig)
		cookie = append(cookie, lcMainCookie)

		csmHitCookieConfig := cookieConfig["csm_hit"].(map[interface{}]interface{})
		csmHitCookie := &http.Cookie{}
		autoSetCookie(csmHitCookie, csmHitCookieConfig)
		cookie = append(cookie, csmHitCookie)
		cookies = append(cookies, cookie)

	}

	return cookies
}

func autoSetCookie(cookie *http.Cookie, cookieConfig map[interface{}]interface{}) {
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
