package main

import (
	"log"

	"github.com/iamthe1whoknocks/rateLimiter/pkg/router"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializating config : %s", err.Error())
	}

	e := router.Run()

	if err := e.Start(":8082"); err != nil {
		log.Fatalf("error starting server : %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("../../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
