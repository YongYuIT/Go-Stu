package main

import (
	"fmt"
	"github.com/spf13/viper"
)

func main() {
	from_env_refix()
}

//从yaml文件中读取配置信息
//https://www.cnblogs.com/cnblogs-wangzhipeng/p/9484460.html
func from_yaml_file() {
	config := viper.New()
	config.SetConfigName("configtx")
	config.SetConfigType("yaml")
	config.AddConfigPath("./test_path_1/")
	config.AddConfigPath("./test_path_2/")
	if err := config.ReadInConfig(); err != nil {
		fmt.Println(err)
		return
	}
	var orgs = config.Get("organizations").([]interface{})
	var org0 = orgs[0].(map[interface{}]interface{})
	fmt.Println((org0["Name"]))
}

//环境变量中读取配置信息
//https://blog.csdn.net/cs380637384/article/details/81217767
func from_env() {
	config := viper.New()
	name := config.Get("YONG_NAME")
	fmt.Println(name)

	config.BindEnv("YONG_NAME")
	name = config.Get("YONG_NAME")
	fmt.Println(name)
}

//$ export YONG_NAME=fucku
//$ echo $YONG_NAME
//$ go run main.go

func from_env_refix() {
	config := viper.New()
	config.AutomaticEnv()
	config.SetEnvPrefix("BEN")
	name := config.Get("NAME")
	fmt.Println(name)
}

//$ export BEN_NAME=aaaa
//$ echo $BEN_NAME
//$ go run main.go
