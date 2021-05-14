package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"thinking_spider/database"
)

type ProdDetailRecord struct {
	gorm.Model
	Asin               string
	Desc1              string `sql:"type:text;"`
	Desc2              string `sql:"type:text;"`
	Desc3              string `sql:"type:text;"`
	Desc4              string `sql:"type:text;"`
	Desc5              string `sql:"type:text;"`
	SoldBy             string
	SoldByAsin         string
	ShipsFrom          string
	DateFirstAvailable string
	ProdDesc           string `sql:"type:text;"`
}

func SaveProdDetailRecord(record *ProdDetailRecord) {
	if !database.CurrentDB.HasTable(record) {
		database.CurrentDB.AutoMigrate(record)
	}
	if database.CurrentDB.NewRecord(record) {
		fmt.Println("create record-->", record.ID)
		database.CurrentDB.Create(record)
	} else {
		fmt.Println("update record-->", record.ID)
		database.CurrentDB.Save(record)
	}
}

type AsinUrl struct {
	DetialUrl string `gorm:"column:detial_url"`
	Asin      string `gorm:"column:asin"`
}

func GetUrlByKeyWords(keyword string) *[]AsinUrl {
	getUrlByAsin := fmt.Sprintf("select detial_url, asin\nfrom (select detial_url, asin, row_number() over (partition by asin order by created_at desc) as aindex\n      from key_word_prod_records_source\n      where key_word = '%s'\n        and asin not in (select asin from prod_detail_records)) t\nwhere t.aindex = 1\n  and detial_url != ''", keyword)
	var asinUrls []AsinUrl
	database.CurrentDB.Raw(getUrlByAsin).Scan(&asinUrls)
	return &asinUrls
}
