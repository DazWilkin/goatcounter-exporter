package collector

import (
	"log/slog"
	"time"

	"github.com/prometheus/client_golang/prometheus"
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

func BuildFQName(name string, logger *slog.Logger) string {
	logger.Info("Creating Metric",
		"name", name,
	)
	return prometheus.BuildFQName(namespace, subsystem, name)
}
