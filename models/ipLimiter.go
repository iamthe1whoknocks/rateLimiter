package models

import (
	"sync"

	"golang.org/x/time/rate"
)

type IPLimiter struct {
	//Subnet map[string]*rate.Limiter
	*rate.Limiter

	sync.RWMutex
}
