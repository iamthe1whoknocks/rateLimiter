package handler

import (
	"fmt"
	"sync"
	"time"

	"github.com/iamthe1whoknocks/rateLimiter/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/time/rate"
)

//Структура для работы с rate limiter
type Handler struct {
	Subnets map[string]*rate.Limiter
	sync.RWMutex
	Mask         string
	RequestLimit int
	TimeToWait   time.Duration
}

//Создание нового экземпляра обработчика
func New(mask string, requestLimit int, time time.Duration) *Handler {
	return &Handler{
		Subnets:      make(map[string]*rate.Limiter),
		Mask:         mask,
		RequestLimit: requestLimit,
		TimeToWait:   time,
	}
}

// Индикация отсутствия превышения лимита
func (h *Handler) Hello(c echo.Context) error {
	return c.String(200, "Hello world\n")

}

//middleware, ограничивающий количество запросов (rate limit) из одной подсети IPv4 к определенному url
func (h *Handler) LimitMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ip := c.RealIP()

		subnet, err := utils.GetSubnetFromIP(ip, h.Mask)
		if err != nil {
			return c.String(echo.ErrBadRequest.Code, err.Error())
		}

		h.RLock()
		limiter, exists := h.Subnets[subnet]
		h.RUnlock()

		if !exists {
			limiter := utils.CreateLimiter(h.TimeToWait, h.RequestLimit)
			h.Lock()
			h.Subnets[subnet] = limiter
			h.Unlock()
			return checkLimit(c, next, limiter)
		}

		return checkLimit(c, next, limiter)

	}
}

//проверка лимита (вынес в отдельную функция для уменьшения дублирования кода)
func checkLimit(c echo.Context, next echo.HandlerFunc, limiter *rate.Limiter) error {
	if !limiter.Allow() {
		return c.String(429, echo.ErrTooManyRequests.Error())

	} else {
		if err := next(c); err != nil {
			c.Error(err)
		}
		return nil
	}
}

//сброс лимита, сбрасывает лимит для ip-адреса, с которого пришел запрос
func (h *Handler) Drop(c echo.Context) error {
	ip := c.RealIP()

	subnet, err := utils.GetSubnetFromIP(ip, h.Mask)
	if err != nil {
		return c.String(echo.ErrBadRequest.Code, err.Error())
	}

	limiter := utils.CreateLimiter(h.TimeToWait, h.RequestLimit)

	h.Lock()
	h.Subnets[subnet] = limiter
	h.Unlock()

	return c.String(200, fmt.Sprintf("Dropped rate limiter for subnet : %s\n", subnet))

}
