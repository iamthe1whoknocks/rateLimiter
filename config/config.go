package config

import (
	"github.com/spf13/viper"
)

//Ининциализация конфигурационного файла
func Init() error {
	viper.AddConfigPath("../../config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()

}
