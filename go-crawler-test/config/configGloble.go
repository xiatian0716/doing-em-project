package config

import "time"

var (
	// 请求频率控制(1000毫秒-1秒1个请求)
	RateLimiter = 1000 * time.Millisecond
)
