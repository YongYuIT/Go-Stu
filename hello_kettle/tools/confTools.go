package tools

import (
	"github.com/beevik/etree"
)

func FillConnWithConf(elm *etree.Element, conf map[string]interface{}) {
	elm.FindElement("server").SetText(conf["host"].(string))
	elm.FindElement("username").SetText(conf["user_name"].(string))
	elm.FindElement("password").SetText(conf["pwd"].(string))
	elm.FindElement("database").SetText(conf["db_name"].(string))
	elm.FindElement("port").SetText(conf["port"].(string))
}
