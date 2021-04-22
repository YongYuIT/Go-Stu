package test

import (
	"charts/database"
	"fmt"
	"testing"
)

func Test_do_list(test *testing.T) {
	var results []map[string]interface{}
	//database.CurrentDB.Table("new_release_prod_records").Limit(100).Find(&results)
	database.CurrentDB.Raw("select asin,id,type1,type2,type3,type4 from new_release_prod_records limit 100").Scan(&results)
	fmt.Println(results)
	keys := []string{}
	for s := range results[0] {
		keys = append(keys, s)
	}
	fmt.Println(keys)
}
