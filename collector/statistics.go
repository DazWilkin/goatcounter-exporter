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
			prometheus.Labels{
				"code": client.Code,
			},
		),
		Hits: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "stats_hits"),
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
	var wg sync.WaitGroup

	// Corresponds to /stats/total
	wg.Add(1)
	go func() {
		defer wg.Done()
		total, err := c.client.Stats().Total()
		if err != nil {
			msg := "unable to collect /stats/total"
			if errResponse, ok := err.(*goatcounter.ErrorResponse); ok {
				log.Printf("%s\n%+v", msg, errResponse)
				return
			}

			log.Printf("%s\n%+v", msg, err)
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
			if errResponse, ok := err.(*goatcounter.ErrorResponse); ok {
				log.Printf("%s\n%+v", msg, errResponse)
				return
			}
			log.Printf("%s\n%+v", msg, err)
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
