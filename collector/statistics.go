package collector

import (
	"log/slog"
	"sync"

	"github.com/DazWilkin/goatcounter-exporter/goatcounter"
	"github.com/prometheus/client_golang/prometheus"
)

// StatisticsCollector collects metrics on GoatCounter's /stats endpoint
type StatisticsCollector struct {
	client *goatcounter.Client
	logger *slog.Logger

	Total *prometheus.Desc
	Hits  *prometheus.Desc
}

// NewStatisticsCollector is a function that returns a new StatisticsCollector
func NewStatisticsCollector(client *goatcounter.Client, logger *slog.Logger) *StatisticsCollector {
	logger.Info("Creating StatisticsCollector")
	return &StatisticsCollector{
		client: client,
		logger: logger,

		Total: prometheus.NewDesc(
			BuildFQName("stats_total", logger),
			"List total pageview counts",
			[]string{},
			prometheus.Labels{
				"code": client.Code,
			},
		),
		Hits: prometheus.NewDesc(
			BuildFQName("stats_hits", logger),
			"pageview and visitor stats",
			[]string{
				"path",
				"day",
			},
			prometheus.Labels{
				"code": client.Code,
			},
		),
	}
}

// Collect implements Prometheus' Collector interface and is used to collect metrics
func (c *StatisticsCollector) Collect(ch chan<- prometheus.Metric) {
	logger := c.logger.With("method", "Collect")

	var wg sync.WaitGroup

	// Corresponds to /stats/total
	logger.Info("Enumerating Statistics Total")
	wg.Add(1)
	go func() {
		defer wg.Done()
		total, err := c.client.Stats().Total()
		if err != nil {
			msg := "unable to collect /stats/total"
			if errResponse, ok := err.(*goatcounter.ErrorResponse); ok {
				logger.Info(msg,
					"err", errResponse,
				)
				return
			}

			logger.Info(msg,
				"error", err,
			)
			return
		}

		value := float64(total.Total)
		logger.Debug("Recording measurement",
			"desc", c.Total,
			"value", value,
		)
		ch <- prometheus.MustNewConstMetric(
			c.Total,
			prometheus.GaugeValue,
			value,
		)
	}()

	// Corresponds to /stats/hits
	logger.Info("Enumerating Statistics Hits")
	wg.Add(1)
	go func() {
		defer wg.Done()
		hits, err := c.client.Stats().Hits()
		if err != nil {
			msg := "unable to collect /stats/hits"
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

		for _, hit := range hits.Hits {
			for _, stat := range hit.Stats {
				value := float64(stat.Daily)
				logger.Debug("Recording measurement",
					"desc", c.Hits,
					"value", value,
				)
				ch <- prometheus.MustNewConstMetric(
					c.Hits,
					prometheus.GaugeValue,
					value,
					[]string{
						hit.Path,
						stat.Day,
					}...,
				)
			}
		}
	}()

	wg.Wait()
}

// Describe implements Prometheus' Collector interface and is used to describe metrics
func (c *StatisticsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.Total
	ch <- c.Hits
}
