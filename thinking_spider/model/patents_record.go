package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"thinking_spider/database"
)

type PationsRecord struct {
	gorm.Model
	Title     string
	Status    string
	PID       string
	Index     int
	KeyWords  string `sql:"type:text;"`
	DetailUrl string `sql:"type:text;"`
	PDFUrl    string `sql:"type:text;"`
	PDFPath   string
	IMGPath   string
}

func SavePationsRecord(record *PationsRecord) {
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
