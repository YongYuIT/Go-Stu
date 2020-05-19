package tools

import (
	"github.com/beevik/etree"
	"strings"
)

func HandleConnInfo(conn_els []*etree.Element, connConfs map[string]interface{}) {
	for _, elm := range conn_els {
		if strings.EqualFold("mysql_tmp", elm.FindElement("name").Text()) {
			conf_info := connConfs["mysql_conn"].(map[string]interface{})
			FillConnWithConf(elm, conf_info)
		} else if strings.EqualFold("pg_tmp", elm.FindElement("name").Text()) {
			conf_info := connConfs["pg_conn"].(map[string]interface{})
			FillConnWithConf(elm, conf_info)
		}
	}
}

func FillConnWithConf(elm *etree.Element, conf map[string]interface{}) {
	elm.FindElement("server").SetText(conf["host"].(string))
	elm.FindElement("username").SetText(conf["user_name"].(string))
	elm.FindElement("password").SetText(conf["pwd"].(string))
	elm.FindElement("database").SetText(conf["db_name"].(string))
	elm.FindElement("port").SetText(conf["port"].(string))
}
