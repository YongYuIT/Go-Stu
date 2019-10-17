package main

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

func main() {
	config := viper.New()
	config.AddConfigPath("/home/yong/Go-Stu20191008001/ReadFabricSourceCode/stu_orderer_load_InitViper/tmp_fabric/")
	config.SetConfigName("orderer")
	config.SetEnvPrefix("ORDERER")
	config.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	config.SetEnvKeyReplacer(replacer)

	if err := config.ReadInConfig(); err != nil {
		fmt.Printf("Error reading configuration: %s", err)
	}
	fmt.Println(config.ConfigFileUsed())
	///////////////////////////////////////////////////////////////////////////////////////
}
