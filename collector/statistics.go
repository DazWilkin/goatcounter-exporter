package collector

import (
	"log"
	"sync"

	"github.com/DazWilkin/goatcounter-exporter/goatcounter"
	"github.com/prometheus/client_golang/prometheus"
)

// StatisticsCollector collects metrics on GoatCounter's /stats endpoint
type StatisticsCollector struct {
	client *goatcounter.Client

	Total *prometheus.Desc
	Hits  *prometheus.Desc
}

// NewStatisticsCollector is a function that returns a new StatisticsCollector
func NewStatisticsCollector(client *goatcounter.Client) *StatisticsCollector {
	return &StatisticsCollector{
		client: client,

		Total: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_total"),
			"List total pageview counts",
			[]string{},
			nil,
		),
		Hits: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_hits"),
			"pageview and visitor stats",
			[]string{
				"path",
				"day",
			},
			nil,
		),
	}
}

// Collect implements Prometheus' Collector interface and is used to collect metrics
func (c *StatisticsCollector) Collect(ch chan<- prometheus.Metric) {
	var wg sync.WaitGroup

	// Corresponds to /stats/total
	wg.Add(1)
	go func() {
		defer wg.Done()
		total, err := c.client.Stats().Total()
		if err != nil {
			msg := "unable to collect /stats/total"
			log.Print(msg)
			return
		}

		ch <- prometheus.MustNewConstMetric(
			c.Total,
			prometheus.GaugeValue,
			float64(total.Total),
		)
	}()

	// Corresponds to /stats/hits
	wg.Add(1)
	go func() {
		defer wg.Done()
		hits, err := c.client.Stats().Hits()
		if err != nil {
			msg := "unable to collect /stats/hits"
			log.Print(msg)
			return
		}

		for _, hit := range hits.Hits {
			for _, stat := range hit.Stats {
				ch <- prometheus.MustNewConstMetric(
					c.Hits,
					prometheus.GaugeValue,
					float64(stat.Daily),
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
