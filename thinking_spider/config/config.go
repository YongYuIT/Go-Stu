package config

import "fmt"
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
		return
	}

	var database = config.Get("database").(map[string]interface{})
	//此处代码用反射精简
	DBConn.Conn = database["conn"].(string)
	DBConn.Type = database["type"].(string)

	var apply = config.Get("apply").(string)
	fmt.Println("apply-->", apply)
	applyConfig := config.Get(apply).(map[string]interface{})
	//此处代码用反射精简
	CurrentDefaultConfig.WebSite = applyConfig["web_site"].(string)
	CurrentDefaultConfig.MaxDeep = applyConfig["max_deep"].(int)
	CurrentDefaultConfig.DelaySpider = applyConfig["delay_spider"].(int)
	CurrentDefaultConfig.KeyWords = applyConfig["key_words"].(string)
	CurrentDefaultConfig.PageHandlerQue = applyConfig["page_handler_que"].(string)
	CurrentDefaultConfig.ProductItemsHandlerQue = applyConfig["product_items_handler_que"].(string)
	CurrentDefaultConfig.PagesKey = applyConfig["pages_key"].(string)
	CurrentDefaultConfig.PageCurrentQue = applyConfig["page_current_que"].(string)
	CurrentDefaultConfig.PageNextQue = applyConfig["page_next_que"].(string)
	CurrentDefaultConfig.PageAttr = applyConfig["page_attr"].(string)
	CurrentDefaultConfig.PageUrlTag = applyConfig["page_url_tag"].(string)
	CurrentDefaultConfig.PageStartTag = applyConfig["page_start_tag"].(string)

	priceLevelConfig := applyConfig["price_level"].(map[string]interface{})
	CurrentDefaultConfig.PriceLevelConfig = &PriceLevelConfig{}
	CurrentDefaultConfig.PriceLevelConfig.PriceStrQue = priceLevelConfig["price_str_que"].(string)
	CurrentDefaultConfig.PriceLevelConfig.PriceListQue = priceLevelConfig["price_list_que"].(string)

	itemsConfig := applyConfig["items"].(map[string]interface{})
	CurrentDefaultConfig.ItemConfig = &ItemConfig{}
	CurrentDefaultConfig.ItemConfig.ProductItemQue = itemsConfig["product_item_que"].(string)

	itemConfig := itemsConfig["item"].(map[string]interface{})
	CurrentDefaultConfig.ItemConfig.Item = &Item{}
	//此处代码用反射精简
	CurrentDefaultConfig.ItemConfig.Item.ItemAsinAttr = itemConfig["item_asin_attr"].(string)
	CurrentDefaultConfig.ItemConfig.Item.ItemUUIDAttr = itemConfig["item_uuid_attr"].(string)
	CurrentDefaultConfig.ItemConfig.Item.ItemDescQue = itemConfig["item_desc_que"].(string)
	CurrentDefaultConfig.ItemConfig.Item.ItemDescAttr = itemConfig["item_desc_attr"].(string)
	CurrentDefaultConfig.ItemConfig.Item.ItemSalesQue = itemConfig["item_sales_que"].(string)
	CurrentDefaultConfig.ItemConfig.Item.ItemIndex = itemConfig["item_index_attr"].(string)
	CurrentDefaultConfig.ItemConfig.Item.ItemPriceQue = itemConfig["item_price_que"].(string)

	fmt.Println("get init config-->", CurrentDefaultConfig)
}
