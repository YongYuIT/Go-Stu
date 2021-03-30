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
	Sales      int
	Price      float32 `sql:"type:decimal(10,2);"`
	ShopName   string
	PriceLevel string
}

func NewKeyWordProdRecord() *KeyWordProdRecord {
	return &KeyWordProdRecord{
		Page:      -1,
		PageIndex: -1,
		Sales:     -1,
		Price:     -1,
	}
}
