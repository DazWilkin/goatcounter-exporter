package collector

import (
	"log/slog"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	namespace string = "goatcounter"
	subsystem string = "exporter"
)

func BuildFQName(name string, logger *slog.Logger) string {
	logger.Info("Creating Metric",
		"name", name,
	)
	return prometheus.BuildFQName(namespace, subsystem, name)
}
