package main

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/spf13/viper"
	"hello_kettle/tools"
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

	connConfs := config.Get("conns").(map[string]interface{})
	mysql_conf_info := connConfs["mysql_conn"].(map[string]interface{})
	mysql_info := tools.GetConnStrForMySql(mysql_conf_info)
	mySqlConn := tools.GetConn("mysql", mysql_info)
	if mySqlConn != nil {
		defer mySqlConn.Close()
	}

	pg_conf_info := connConfs["pg_conn"].(map[string]interface{})
	pg_info := tools.GetConnStrForPG(pg_conf_info)
	pgConn := tools.GetConn("postgres", pg_info)
	if pgConn != nil {
		defer pgConn.Close()
	}

	tools.HandleFullFile("duichou_tmp_mysql_gp_full.ktr", config, mySqlConn)
	tools.HandleMaxTimeFile("duichou_tmp_mysql_gp_maxTime.ktr", config)
	tools.HandleIncreFile("duichou_tmp_mysql_gp_incre.ktr", config, mySqlConn)
}
