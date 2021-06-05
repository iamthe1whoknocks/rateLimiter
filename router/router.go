package router

import (
	"github.com/iamthe1whoknocks/rateLimiter/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Создание роутера
func New(h *handler.Handler) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//роут с включенным rate limiter
	e.GET("/limit", h.Hello, h.LimitMiddleware)

	//для сброса лимита (сброс по ip-адресу из запроса)
	e.GET("/drop", h.Drop)

	return e
}
