package main

import (
	"log"
	"time"

	"github.com/iamthe1whoknocks/rateLimiter/handler"
	"github.com/iamthe1whoknocks/rateLimiter/router"

	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatal(err.Error())
	}

	mask := viper.GetString("mask")
	requestLimit := viper.GetInt("request_limit")
	t := viper.GetInt("time")
	timeToWait := time.Duration(t) * time.Minute

	h := handler.New(mask, requestLimit, timeToWait)

	e := router.New(h)

	e.Start(":8083")

}

func initConfig() error {
	viper.AddConfigPath("../../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
