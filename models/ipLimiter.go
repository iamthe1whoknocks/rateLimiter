package models

import (
	"sync"

	"gopkg.in/labstack/echo.v4"

	"golang.org/x/time/rate"
)

type IPLimiter struct {
	//Subnet map[string]*rate.Limiter
	*rate.Limiter

	sync.RWMutex
}

func (l *IPLimiter) LimitMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		if !l.Allow() {
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
