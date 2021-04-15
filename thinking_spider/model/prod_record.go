package model

import "github.com/jinzhu/gorm"

type ProdRecord struct {
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
