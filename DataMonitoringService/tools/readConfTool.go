package tools

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type TabMonItem struct {
	Tabname  string
	ScheName string
	DBConf   *DBConfig
}

func GetZeroTabMonConf() ([]TabMonItem, error) {
	conf, err := getConfig("TabDataIncrementMon", "./config/")
	if err != nil {
		return nil, err
	}
	tabs := conf.Get("zero_incre_tab_mon").([]interface{})
	tabinfos := make([]TabMonItem, 0)
	for _, v := range tabs {
		vinfo := strings.Split(v.(string), ".")
		dbConfig := GetDBConfigByID(vinfo[0])
		if vinfo[2] != "*" {
			tabinfos = append(tabinfos, TabMonItem{vinfo[2], vinfo[1], dbConfig})
			fmt.Println("GetZeroTabMonConf", v)
		} else {
			//需补充sch.*的情况
			sDbTool := (&SchemaTabInfoDBTool{}).InitTool(dbConfig.ID).(*SchemaTabInfoDBTool)
			tables := sDbTool.ReadAllTabsUnderSchema(vinfo[1])
			if tables != nil {
				for _, v1 := range tables {
					tabinfos = append(tabinfos, TabMonItem{v1.TabName, vinfo[1], dbConfig})
				}
			}
		}
	}
	return tabinfos, nil
}

type DBConfig struct {
	ID       string
	IPPort   string
	UserName string
	Passwd   string
	DBName   string
}

func GetDBConfig() ([]DBConfig, error) {
	conf, err := getConfig("TabDataIncrementMon", "./config/")
	if err != nil {
		return nil, err
	}
	dbs := conf.Get("db_conn").([]interface{})
	fmt.Println("dbs", dbs)
	dbInfos := make([]DBConfig, 0)
	for _, v := range dbs {
		var db DBConfig
		vv := v.(map[interface{}]interface{})
		fmt.Println(vv)
		db.DBName = vv["db_name"].(string)
		db.IPPort = vv["ip_port"].(string)
		db.UserName = vv["user_name"].(string)
		db.Passwd = vv["pwd"].(string)
		db.ID = vv["id"].(string)
		dbInfos = append(dbInfos, db)
	}
	return dbInfos, nil
}

func GetDBConfigByID(id string) *DBConfig {
	conf, err := getConfig("TabDataIncrementMon", "./config/")
	if err != nil {
		return nil
	}
	dbs := conf.Get("db_conn").([]interface{})
	fmt.Println("dbs", dbs)
	var dbInfo *DBConfig = nil
	for _, v := range dbs {
		var db DBConfig
		vv := v.(map[interface{}]interface{})
		fmt.Println(vv)
		db.DBName = vv["db_name"].(string)
		db.IPPort = vv["ip_port"].(string)
		db.UserName = vv["user_name"].(string)
		db.Passwd = vv["pwd"].(string)
		db.ID = vv["id"].(string)
		if strings.EqualFold(db.ID, id) {
			dbInfo = &db
			break
		}
	}
	return dbInfo
}

type SendEmailConf struct {
	Username string
	Passwd   string
	STMPHost string
	STMPPort string
}

func getConfig(conf_name string, conf_path string) (*viper.Viper, error) {
	config := viper.New()
	config.SetConfigName(conf_name)
	config.SetConfigType("yaml")
	config.AddConfigPath(conf_path)
	if err := config.ReadInConfig(); err != nil {
		fmt.Println(err)
		return nil, err
	}
	return config, nil
}
