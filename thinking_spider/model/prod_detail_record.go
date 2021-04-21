package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"thinking_spider/database"
)

type ProdDetailRecord struct {
	gorm.Model
	Asin       string
	Desc1      string `sql:"type:text;"`
	Desc2      string `sql:"type:text;"`
	Desc3      string `sql:"type:text;"`
	Desc4      string `sql:"type:text;"`
	Desc5      string `sql:"type:text;"`
	SoldBy     string
	SoldByAsin string
	ShipsFrom  string
	ProdDesc   string `sql:"type:text;"`
}

func SaveProdDetailRecord(record *ProdDetailRecord) {
	if !database.CurrentDB.HasTable(record) {
		database.CurrentDB.AutoMigrate(record)
	}
	database.CurrentDB.Create(record)
}

type AsinUrl struct {
	DetialUrl string `gorm:"column:detial_url"`
	Asin      string `gorm:"column:asin"`
}

func GetUrlByKeyWords(keyword string) *[]AsinUrl {
	getUrlByAsin := fmt.Sprintf("select detial_url, asin from (select detial_url, asin, row_number() over (partition by asin order by created_at desc) as aindex from key_word_prod_records where key_word='%s') t where t.aindex = 1 and detial_url != ''", keyword)
	var asinUrls []AsinUrl
	database.CurrentDB.Raw(getUrlByAsin).Scan(&asinUrls)
	return &asinUrls
}
