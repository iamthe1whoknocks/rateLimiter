package router

import (
	"github.com/iamthe1whoknocks/rateLimiter/handler"
	"github.com/labstack/echo"
)

func New(h *handler.Handler) *echo.Echo {

	e := echo.New()
	e.Use(h.LimitMiddleware)
	e.GET("/", h.Hello)

	return e
}
