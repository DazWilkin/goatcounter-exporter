package collector

import (
	"log"

	"github.com/DazWilkin/goatcounter-exporter/goatcounter"
	"github.com/prometheus/client_golang/prometheus"
)

// PathsCollector collects metrics on GoatCounter's /paths endpoint
type PathsCollector struct {
	client *goatcounter.Client

	Total *prometheus.Desc
}

// NewPathsCollector is a function that returns a new PathsCollector
func NewPathsCollector(client *goatcounter.Client) *PathsCollector {
	return &PathsCollector{
		client: client,

		Total: prometheus.NewDesc(
			prometheus.BuildFQName(namespace, subsystem, "paths_total"),
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
	pathsResponse, err := c.client.Paths().List()
	if err != nil {
		msg := "unable to collect /paths"
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
		float64(len(pathsResponse.Paths)),
	)
}

// Describe is a method that implements Prometheus' Collector interface and is used to describe metrics
func (c *PathsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.Total
}
