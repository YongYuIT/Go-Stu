package config

import "fmt"
import "github.com/spf13/viper"

type SprierConfig struct {
	WebSite                string
	MaxDeep                int
	KeyWords               string
	PageHandlerQue         string
	ProductItemsHandlerQue string
	PagesKey               string
	PageCurrentQue         string
	PageNextQue            string
	PageAttr               string
	PageUrlTag             string
	PageStartTag           string
	Items                  *Items
}

type Items struct {
	ProductItemQue string
	Item           *Item
}

type Item struct {
	ItemAsinAttr string
	ItemUUIDAttr string
	ItemDescQue  string
	ItemDescAttr string
	ItemSalesQue string
	ItemIndex    string
}

type DBConfig struct {
	Conn string
	Type string
}

var CurrentSprierConfig = &SprierConfig{}
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
	CurrentSprierConfig.WebSite = applyConfig["web_site"].(string)
	CurrentSprierConfig.MaxDeep = applyConfig["max_deep"].(int)
	CurrentSprierConfig.KeyWords = applyConfig["key_words"].(string)
	CurrentSprierConfig.PageHandlerQue = applyConfig["page_handler_que"].(string)
	CurrentSprierConfig.ProductItemsHandlerQue = applyConfig["product_items_handler_que"].(string)
	CurrentSprierConfig.PagesKey = applyConfig["pages_key"].(string)
	CurrentSprierConfig.PageCurrentQue = applyConfig["page_current_que"].(string)
	CurrentSprierConfig.PageNextQue = applyConfig["page_next_que"].(string)
	CurrentSprierConfig.PageAttr = applyConfig["page_attr"].(string)
	CurrentSprierConfig.PageUrlTag = applyConfig["page_url_tag"].(string)
	CurrentSprierConfig.PageStartTag = applyConfig["page_start_tag"].(string)

	itemsConfig := applyConfig["items"].(map[string]interface{})
	CurrentSprierConfig.Items = &Items{}
	CurrentSprierConfig.Items.ProductItemQue = itemsConfig["product_item_que"].(string)

	itemConfig := itemsConfig["item"].(map[string]interface{})
	CurrentSprierConfig.Items.Item = &Item{}
	//此处代码用反射精简
	CurrentSprierConfig.Items.Item.ItemAsinAttr = itemConfig["item_asin_attr"].(string)
	CurrentSprierConfig.Items.Item.ItemUUIDAttr = itemConfig["item_uuid_attr"].(string)
	CurrentSprierConfig.Items.Item.ItemDescQue = itemConfig["item_desc_que"].(string)
	CurrentSprierConfig.Items.Item.ItemDescAttr = itemConfig["item_desc_attr"].(string)
	CurrentSprierConfig.Items.Item.ItemSalesQue = itemConfig["item_sales_que"].(string)
	CurrentSprierConfig.Items.Item.ItemIndex = itemConfig["item_index_attr"].(string)

	fmt.Println("get init config-->", CurrentSprierConfig)
}
