package config

import "net/http"

type SpiderConfig struct {
	Model                  string            `config:"model"`
	TaskIndex              int               `config:"task_index"`
	WebSite                string            `config:"web_site"`
	NewRelease             string            `config:"new_release"`
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
	PricesLevels           string            `config:"prices_levels"`
	RobortQue              string            `config:"robort_que"`
	ItemsConfig            *ItemsConfig      `config:"items"`
	PriceLevelConfig       *PriceLevelConfig `config:"price_level"`
	Cookies                []*http.Cookie    `cookies`
	DetailsConfig          *DetailsConfig    `config:"details"`
}

type ItemsConfig struct {
	ProductItemQue string `config:"product_item_que"`
	Item           *Item  `config:"item"`
}

type PriceLevelConfig struct {
	PriceStrQue  string `config:"price_str_que"`
	PriceListQue string `config:"price_list_que"`
}

type DetailsConfig struct {
	DescsQue    string `config:"descs_que"`
	AsinQue     string `config:"asin_que"`
	SoldByQue   string `config:"sold_by_que"`
	SoldIdQue   string `config:"sold_id_que"`
	ProdDescQue string `config:"prod_desc_que"`
}

type Item struct {
	ItemAsinAttr     string `config:"item_asin_attr"`
	ItemUUIDAttr     string `config:"item_uuid_attr"`
	ItemDescQue      string `config:"item_desc_que"`
	ItemPriceQue     string `config:"item_price_que"`
	ItemDescAttr     string `config:"item_desc_attr"`
	ItemRangeQue     string `config:"item_range_que"`
	ItemIndexAttr    string `config:"item_index_attr"`
	ItemStartsQue    string `config:"item_starts_que"`
	ItemDetailUrlQue string `config:"item_detail_url_que"`
	ItemImgUrlQue    string `config:"item_img_url_que"`
}
