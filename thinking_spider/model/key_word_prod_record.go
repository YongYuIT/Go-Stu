package model

import (
	"github.com/jinzhu/gorm"
)

type KeyWordProdRecord struct {
	gorm.Model
	KeyWord    string
	Page       int
	PageIndex  int
	Asin       string
	Uuid       string
	Desc       string
	Range      int
	Price      float32 `sql:"type:decimal(10,2);"`
	ShopName   string
	PriceLevel string
	DeliverTo  string
}

func NewKeyWordProdRecord() *KeyWordProdRecord {
	return &KeyWordProdRecord{
		Page:      -1,
		PageIndex: -1,
		Range:     -1,
		Price:     -1,
	}
}
