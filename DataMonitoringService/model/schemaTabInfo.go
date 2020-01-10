package model

type SchemaTabInfo struct {
	TabCatalog string `gorm:"column:table_catalog"`
	TabSchema  string `gorm:"column:table_schema"`
	TabName    string `gorm:"column:table_name"`
	TabType    string `gorm:"column:table_type"`
}
