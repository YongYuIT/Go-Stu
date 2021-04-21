package model

type BasicProdInfo struct {
	Asin       string
	Titles     string
	Ratings    int
	Starts     float32 `sql:"type:decimal(10,2);"`
	DeliverTo  string
	DetialUrl  string `sql:"type:text;"`
	MainPicUrl string `sql:"type:text;"`
}
