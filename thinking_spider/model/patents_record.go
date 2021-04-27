package model

import (
	"github.com/jinzhu/gorm"
	"thinking_spider/database"
)

type PationsRecord struct {
	gorm.Model
	Title     string
	Status    string
	PID       string
	Index     int
	KeyWords  string
	DetailUrl string
	PDFUrl    string
	PDFPath   string
	IMGPath   string
}

func SavePationsRecord(record *PationsRecord) {
	if !database.CurrentDB.HasTable(record) {
		database.CurrentDB.AutoMigrate(record)
	}
	database.CurrentDB.Create(record)
}
