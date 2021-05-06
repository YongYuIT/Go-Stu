package model

import "github.com/jinzhu/gorm"
import "thinking_spider/database"

type NewReleaseProdRecord struct {
	gorm.Model
	BasicProdInfo
	Type1     string
	Type2     string
	Type3     string
	Type4     string
	Type5     string
	Type6     string
	Price     string
	Index     int
	TaskIndex int
}

func NewNewReleaseProdRecord() *NewReleaseProdRecord {
	release := &NewReleaseProdRecord{
		Index: -1,
	}
	release.Ratings = -1
	return release
}

func SaveNewReleaseProdRecord(record *NewReleaseProdRecord) {
	if !database.CurrentDB.HasTable(record) {
		database.CurrentDB.AutoMigrate(record)
	}
	database.CurrentDB.Create(record)
}
