package main

import (
	"fmt"
	"github.com/beevik/etree"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"hello_kettle/tools"
	"strings"
)

func main() {
	config := viper.New()
	config.SetConfigName("conf")
	config.SetConfigType("yaml")
	config.AddConfigPath("./")
	if err := config.ReadInConfig(); err != nil {
		fmt.Println(err)
		return
	}

	doc := etree.NewDocument()
	err := doc.ReadFromFile("duichou_tmp_mysql_gp_full.ktr")
	if err != nil {
		fmt.Println("read doc err-->", err)
		return
	}
	//获取连接信息
	ktr_els := doc.SelectElement("transformation")
	if ktr_els == nil {
		fmt.Println("not found ktr els!")
		return
	}
	conn_els := ktr_els.FindElements("connection")
	fmt.Println("find conn-->", len(conn_els))

	connConfs := config.Get("conns").(map[string]interface{})

	mysql_info := ""
	psql_info := ""
	for _, elm := range conn_els {
		if strings.EqualFold("mysql_tmp", elm.FindElement("name").Text()) {
			conf_info := connConfs["mysql_conn"].(map[string]interface{})
			tools.FillConnWithConf(elm, conf_info)
			mysql_info = fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
				elm.FindElement("username").Text(),
				elm.FindElement("password").Text(),
				elm.FindElement("server").Text(),
				elm.FindElement("port").Text(),
				elm.FindElement("database").Text())
		} else if strings.EqualFold("pg_tmp", elm.FindElement("name").Text()) {
			conf_info := connConfs["pg_conn"].(map[string]interface{})
			tools.FillConnWithConf(elm, conf_info)
			psql_info = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
				elm.FindElement("server").Text(),
				elm.FindElement("port").Text(),
				elm.FindElement("username").Text(),
				elm.FindElement("password").Text(),
				elm.FindElement("database").Text())
		}
	}
	fmt.Println("mysql conn str-->", mysql_info)
	fmt.Println("pg conn str-->", psql_info)

	//获取表信息
	mySqlConn, err := gorm.Open("mysql", mysql_info)
	if err != nil {
		fmt.Println("mysql conn failed-->", err)
		return
	}
	fmt.Println("conn mysql success")
	defer mySqlConn.Close()

	pgConn, err := gorm.Open("postgres", psql_info)
	if err != nil {
		fmt.Println("postgres conn failed-->", err)
		return
	}
	fmt.Println("postgres mysql success")
	defer pgConn.Close()

	tabs := config.Get("need_hand_tab").([]interface{})
	for _, tab := range tabs {
		fmt.Println("need tab-->", tab)
		file := tools.HandleTab(ktr_els, tab.(string), mySqlConn, nil)
		doc.WriteToFile(file + "_full.ktr")
	}
}
