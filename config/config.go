package config

import (
	"github.com/spf13/viper"
	"fmt"
)

func Init(fileName string) *viper.Viper {
	v := viper.New()
	v.SetConfigName(fileName)
	v.SetConfigType("yaml")
	v.AddConfigPath("config/")
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("file not found")
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
		}
	}
	return v
}

func InitWithRootPath() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("../../config/")
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("file not found")
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
			fmt.Println("File found")
		}
	}
	return v
}