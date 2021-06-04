package handler

import (
	"sync"
	"time"

	"github.com/iamthe1whoknocks/rateLimiter/utils"
	"github.com/labstack/echo"
	"golang.org/x/time/rate"
)

type Handler struct {
	Subnets map[string]*rate.Limiter
	sync.RWMutex
	Mask         string
	RequestLimit int
	TimeToWait   time.Duration
}

func New(mask string, requestLimit int, time time.Duration) *Handler {
	return &Handler{
		Subnets:      make(map[string]*rate.Limiter),
		Mask:         mask,
		RequestLimit: requestLimit,
		TimeToWait:   time,
	}
}

func (h *Handler) Hello(c echo.Context) error {
	c.String(200, "Hello world")
	return nil

}

func (h *Handler) LimitMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ip := c.RealIP()

		subnet, err := utils.GetSubnetFromIP(ip, h.Mask)
		if err != nil {
			return c.String(echo.ErrBadRequest.Code, err.Error())
		}

		limiter := h.getSubnet(subnet)

		if !limiter.Allow() {
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

//Получение ограничителя для выбранной подсети
func (h *Handler) getSubnet(subnet string) *rate.Limiter {
	h.RLock()
	limiter, exists := h.Subnets[subnet]
	h.RUnlock()
	if !exists {
		limiter = h.addSubnet(subnet)
		return limiter
	}
	return limiter
}

//Добавление подсети в список
func (h *Handler) addSubnet(subnet string) *rate.Limiter {
	rt := rate.Every(h.TimeToWait * time.Minute)
	limiter := rate.NewLimiter(rt, h.RequestLimit)
	h.Lock()
	h.Subnets[subnet] = limiter
	h.Unlock()
	return limiter
}
