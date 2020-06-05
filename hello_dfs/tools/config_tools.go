package tools

import (
	"fmt"
	v "github.com/spf13/viper"
)

func GetConfig() (*v.Viper, error) {
	config := v.New()
	config.SetConfigName("conf")
	config.SetConfigType("yaml")
	config.AddConfigPath("./")
	if err := config.ReadInConfig(); err != nil {
		fmt.Println("read config err-->", err)
		return nil, err
	}
	return config, nil
}
