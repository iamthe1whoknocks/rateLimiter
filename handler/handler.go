package handler

import (
	"github.com/iamthe1whoknocks/rateLimiter/models"
	"gopkg.in/labstack/echo.v4"
)

type Handler struct {
	Limiter *models.IPLimiter
}

func (h *Handler) Hello(c echo.Context) error {

	c.String(200, "Hello world")
	return nil

}

func (h *Handler) LimitMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		if !h.Limiter.Allow() {
			c.String(429, echo.ErrTooManyRequests.Error())
			return nil
		} else {
			if err := next(c); err != nil {
				c.Error(err)
			}
			return nil
		}

	}
}
