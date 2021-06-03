package main

import (
	"time"

	"github.com/iamthe1whoknocks/rateLimiter/handler"
	"github.com/iamthe1whoknocks/rateLimiter/models"
	"golang.org/x/time/rate"

	"github.com/spf13/viper"
	"gopkg.in/labstack/echo.v4"
)

func main() {
	e := echo.New()

	rt := rate.Every(1 * time.Minute)

	limiter := rate.NewLimiter(rt, 5)
	ipLimiter := &models.IPLimiter{Limiter: limiter}

	h := &handler.Handler{Limiter: ipLimiter}

	e.GET("/", h.Hello, h.Limiter.LimitMiddleware)

	e.Start(":8083")

}

func initConfig() error {
	viper.AddConfigPath("../../configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
