package config

type SpiderConfig struct {
	WebSite                string
	MaxDeep                int
	DelaySpider            int
	KeyWords               string
	PageHandlerQue         string
	ProductItemsHandlerQue string
	PagesKey               string
	PageCurrentQue         string
	PageNextQue            string
	PageAttr               string
	PageUrlTag             string
	PageStartTag           string
	CurrentPriceLevel      string //not read from yaml file
	ItemConfig             *ItemConfig
	PriceLevelConfig       *PriceLevelConfig
}

type ItemConfig struct {
	ProductItemQue string
	Item           *Item
}

type PriceLevelConfig struct {
	PriceStrQue  string
	PriceListQue string
}

type Item struct {
	ItemAsinAttr string
	ItemUUIDAttr string
	ItemDescQue  string
	ItemPriceQue string
	ItemDescAttr string
	ItemSalesQue string
	ItemIndex    string
}
