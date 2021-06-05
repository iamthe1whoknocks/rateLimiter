package utils

import (
	"net"
	"strings"
	"time"

	"golang.org/x/time/rate"
)

//Получение подсети из ip-адреса, полученного из запроса, и маски подсети.
func GetSubnetFromIP(ip string, mask string) (string, error) {
	var address strings.Builder
	address.WriteString(ip)
	address.WriteString("/")
	address.WriteString(mask)
	_, subnet, err := net.ParseCIDR(address.String())
	if err != nil {
		return "", err
	}
	return subnet.String(), nil
}

// Создание rate limiter из конфигурационных данных
func CreateLimiter(timeToWait time.Duration, requestLimit int) (limiter *rate.Limiter) {
	rt := rate.Every(timeToWait * time.Minute / time.Duration(requestLimit))
	limiter = rate.NewLimiter(rt, requestLimit)
	return limiter
}
