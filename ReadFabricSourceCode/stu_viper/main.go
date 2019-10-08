package main

import (
	"fmt"
	"github.com/spf13/viper"
	"reflect"
)

//https://www.cnblogs.com/cnblogs-wangzhipeng/p/9484460.html
func main() {
	from_env()
}

//从yaml文件中读取配置信息
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
	var orgs = config.Get("organizations").([]interface{}) //Go语言的类型转换和类型断言
	var org0 = orgs[0].(map[interface{}]interface{})
	fmt.Println((org0["Name"]))
}

//环境变量中读取配置信息
func from_env() {
	config := viper.New()
	config.AutomaticEnv()
	fmt.Println(reflect.TypeOf(config))
}
