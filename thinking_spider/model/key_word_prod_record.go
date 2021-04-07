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
	Titles     string
	Ratings    int
	Starts     float32 `sql:"type:decimal(10,2);"`
	Price      float32 `sql:"type:decimal(10,2);"`
	ShopName   string
	PriceLevel string
	DeliverTo  string
	DetialUrl  string `sql:"type:text;"`
}

func NewKeyWordProdRecord() *KeyWordProdRecord {
	return &KeyWordProdRecord{
		Page:      -1,
		PageIndex: -1,
		Ratings:   -1,
		Price:     -1,
	}
}
