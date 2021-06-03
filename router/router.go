package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {

	e := echo.New()
	e.Use(middleware.RateLimiterWithConfig(&middleware.RateLimiterConfig{
		
	})

	return e
}
