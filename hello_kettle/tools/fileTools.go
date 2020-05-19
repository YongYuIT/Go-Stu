package tools

import (
	"fmt"
	"github.com/beevik/etree"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"strings"
)

func HandleFullFile(file_name string, config *viper.Viper, mySqlConn *gorm.DB) {
	doc := etree.NewDocument()
	err := doc.ReadFromFile(file_name)
	if err != nil {
		fmt.Println("read doc err-->", err)
		return
	}
	ktr_els := doc.SelectElement("transformation")

	//处理连接信息
	HandleFileConnInfo(ktr_els, config)

	tabs := config.Get("need_hand_tabs").([]interface{})
	for _, tab := range tabs {
		fmt.Println("need tab-->", tab)
		tab_info := tab.(map[interface{}]interface{})
		file := HandleTabFullMsg(ktr_els, tab_info, mySqlConn)
		doc.WriteToFile(file + "_full.ktr")
	}
}
func HandleIncreFile(file_name string, config *viper.Viper, mySqlConn *gorm.DB) {
	doc := etree.NewDocument()
	err := doc.ReadFromFile(file_name)
	if err != nil {
		fmt.Println("read doc err-->", err)
		return
	}
	ktr_els := doc.SelectElement("transformation")

	//处理连接信息
	HandleFileConnInfo(ktr_els, config)
	tabs := config.Get("need_hand_tabs").([]interface{})
	for _, tab := range tabs {
		fmt.Println("need tab-->", tab)
		tab_info := tab.(map[interface{}]interface{})
		file := HandleTabIncreMsg(ktr_els, tab_info, mySqlConn)
		doc.WriteToFile(file + "_incre.ktr")
	}
}

func HandleMaxTimeFile(file_name string, config *viper.Viper) {
	doc := etree.NewDocument()
	err := doc.ReadFromFile(file_name)
	if err != nil {
		fmt.Println("read doc err-->", err)
		return
	}
	ktr_els := doc.SelectElement("transformation")

	//处理连接信息
	HandleFileConnInfo(ktr_els, config)

	tabs := config.Get("need_hand_tabs").([]interface{})
	for _, tab := range tabs {
		tab_info := tab.(map[interface{}]interface{})
		mysqlTabName := strings.Split(tab_info["source_tab"].(string), ".")[1]
		maxStr := "select case when max(update_time) is null then '1900-01-01 00:00:01' else  to_char(max(update_time),'YYYY-MM-DD HH24:MI:SS') end as max_MODIFYRQ from " + mysqlTabName
		steps := ktr_els.FindElements("step")
		for i := 0; i < len(steps); i++ {
			if strings.EqualFold("TabInput", steps[i].FindElement("name").Text()) {
				steps[i].FindElement("sql").SetText(maxStr)
			}
		}
		doc.WriteToFile(mysqlTabName + "_max_time.ktr")
	}
}

//处理连接信息
func HandleFileConnInfo(ktr_els *etree.Element, config *viper.Viper) {
	if ktr_els == nil {
		fmt.Println("not found ktr els!")
		return
	}
	conn_els := ktr_els.FindElements("connection")
	fmt.Println("find conn-->", len(conn_els))
	connConfs := config.Get("conns").(map[string]interface{})
	HandleConnInfo(conn_els, connConfs)
}
