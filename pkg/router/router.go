package router

import (
	"github.com/labstack/echo"
)

func Run() *echo.Echo {
	e := echo.New()

	e.GET("/", helloWorld)

	return e

}

func helloWorld(c echo.Context) error {
	return c.String(200, "hello world")
}
