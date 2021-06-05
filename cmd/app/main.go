package main

import (
	"log"
	"time"

	"github.com/iamthe1whoknocks/rateLimiter/config"
	"github.com/iamthe1whoknocks/rateLimiter/handler"
	"github.com/iamthe1whoknocks/rateLimiter/router"

	"github.com/spf13/viper"
)

func main() {

	//инициализация конфигурационного файла
	if err := config.Init(); err != nil {
		log.Fatal(err.Error())
	}

	//получение данных из  конфигурационного файла
	mask := viper.GetString("mask")
	requestLimit := viper.GetInt("request_limit")
	timeToWait := time.Duration(viper.GetInt64("time"))

	//Создание экземпляра handler
	h := handler.New(mask, requestLimit, timeToWait)

	//создание роутера
	e := router.New(h)

	e.Logger.Fatal(e.Start(":8083"))

}
