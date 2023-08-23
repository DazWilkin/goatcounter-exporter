package collector

import (
	"time"

	"golang.org/x/time/rate"
)

const (
	namespace string = "goatcounter"
	subsystem string = "exporter"
)

var (
	// GoatCounter rate limit is 4qps
	// https://pretired.goatcounter.com/help/api#rate-limit-213
	// Rate Limiter is shared across collectors
	ratelimiter = rate.NewLimiter(rate.Every(time.Second), 4)
)
