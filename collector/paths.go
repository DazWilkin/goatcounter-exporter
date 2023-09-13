package collector

import (
	"log/slog"

	"github.com/DazWilkin/goatcounter-exporter/goatcounter"
	"github.com/prometheus/client_golang/prometheus"
)

// PathsCollector collects metrics on GoatCounter's /paths endpoint
type PathsCollector struct {
	client *goatcounter.Client
	logger *slog.Logger

	Total *prometheus.Desc
}

// NewPathsCollector is a function that returns a new PathsCollector
func NewPathsCollector(client *goatcounter.Client, logger *slog.Logger) *PathsCollector {
	logger.Info("Creating PathsCollector")
	return &PathsCollector{
		client: client,
		logger: logger,

		Total: prometheus.NewDesc(
			BuildFQName("paths_total", logger),
			"List total of paths",
			[]string{},
			prometheus.Labels{
				"code": client.Code,
			},
		),
	}
}

// Collect is a method that implements Prometheus' Collector itnerface and is used to collect metrics
func (c *PathsCollector) Collect(ch chan<- prometheus.Metric) {
	logger := c.logger.With("method", "Collect")
	pathsResponse, err := c.client.Paths().List()
	if err != nil {
		msg := "unable to collect /paths"
		if errResponse, ok := err.(*goatcounter.ErrorResponse); ok {
			logger.Info(msg,
				"error", errResponse,
			)
			return

		}
		logger.Info(msg,
			"error", err,
		)
		return
	}

	value := float64(len(pathsResponse.Paths))
	logger.Debug("Recording measurement",
		"desc", c.Total,
		"value", value,
	)
	ch <- prometheus.MustNewConstMetric(
		c.Total,
		prometheus.GaugeValue,
		value,
	)
}

// Describe is a method that implements Prometheus' Collector interface and is used to describe metrics
func (c *PathsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.Total
}
