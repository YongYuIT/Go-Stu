package config

import "net/http"

type SpiderConfig struct {
	Model                  string            `config:"model"`
	WebSite                string            `config:"web_site"`
	MaxDeep                int               `config:"max_deep"`
	DelaySpider            int               `config:"delay_spider"`
	KeyWords               string            `config:"key_words"`
	PageHandlerQue         string            `config:"page_handler_que"`
	ProductItemsHandlerQue string            `config:"product_items_handler_que"`
	PagesKey               string            `config:"pages_key"`
	PageCurrentQue         string            `config:"page_current_que"`
	PageNextQue            string            `config:"page_next_que"`
	PageAttr               string            `config:"page_attr"`
	PageUrlTag             string            `config:"page_url_tag"`
	PageStartTag           string            `config:"page_start_tag"`
	RegionQue              string            `config:"region_que"`
	CurrentPriceLevel      string            //not read from yaml file
	ItemsConfig            *ItemsConfig      `config:"items"`
	PriceLevelConfig       *PriceLevelConfig `config:"price_level"`
	Cookies                []*http.Cookie    `cookies`
}

type ItemsConfig struct {
	ProductItemQue string `config:"product_item_que"`
	Item           *Item  `config:"item"`
}

type PriceLevelConfig struct {
	PriceStrQue  string `config:"price_str_que"`
	PriceListQue string `config:"price_list_que"`
}

type Item struct {
	ItemAsinAttr  string `config:"item_asin_attr"`
	ItemUUIDAttr  string `config:"item_uuid_attr"`
	ItemDescQue   string `config:"item_desc_que"`
	ItemPriceQue  string `config:"item_price_que"`
	ItemDescAttr  string `config:"item_desc_attr"`
	ItemSalesQue  string `config:"item_sales_que"`
	ItemIndexAttr string `config:"item_index_attr"`
}
