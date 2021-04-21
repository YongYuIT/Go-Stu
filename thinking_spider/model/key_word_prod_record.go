package model

import (
	"github.com/jinzhu/gorm"
	"thinking_spider/database"
)

type KeyWordProdRecord struct {
	gorm.Model
	BasicProdInfo
	Uuid       string
	KeyWord    string
	Page       int
	PageIndex  int
	ShopName   string
	PriceLevel string
	Price      float32 `sql:"type:decimal(10,2);"`
}

func NewKeyWordProdRecord() *KeyWordProdRecord {
	kw := &KeyWordProdRecord{
		Page:      -1,
		PageIndex: -1,
		Price:     -1,
	}
	kw.Ratings = -1
	return kw
}

func SaveKeyWordProdRecord(record *KeyWordProdRecord) {
	if !database.CurrentDB.HasTable(record) {
		database.CurrentDB.AutoMigrate(record)
	}
	database.CurrentDB.Create(record)
}
