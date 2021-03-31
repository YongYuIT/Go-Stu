package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)
import "github.com/spf13/viper"

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

	var apply = config.Get("apply").(string)
	fmt.Println("apply-->", apply)
	applyConfig := config.Get(apply).(map[string]interface{})
	autoSetValueFromConfig(CurrentDefaultConfig, applyConfig)

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
