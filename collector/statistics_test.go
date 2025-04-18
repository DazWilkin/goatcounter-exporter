package collector

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"text/template"
	"time"

	"github.com/DazWilkin/goatcounter-exporter/goatcounter"

	"github.com/prometheus/client_golang/prometheus/testutil"
)

const (
	// prometheusStatistics is a template to improve replacement of dynamic data e.g. Code, Instance
	prometheusStatistics string = `
# HELP goatcounter_exporter_stats_total List total pageview counts
# TYPE goatcounter_exporter_stats_total gauge
goatcounter_exporter_stats_total{code="{{ .Code }}",instance="{{ .Instance }}"} 1
# HELP goatcounter_exporter_stats_hits pageview and visitor stats
# TYPE goatcounter_exporter_stats_hits gauge
goatcounter_exporter_stats_hits{code="{{ .Code }}",day="{{ .Day }}",instance="{{ .Instance }}",path="/"} 1
`
)

// TestStatisticsCollector tests StatisticsCollector.Collect
// Collect covers 2 APIs (/stats/total; /stats/hits) represented here by Handler closures
func TestStatisticsCollector(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	code := "test"
	instance := server.URL

	client := goatcounter.NewTestClient(code, instance, logger)
	path := fmt.Sprintf("/%s", goatcounter.Version)

	day := time.Now().Format("2006-01-02")

	mux.HandleFunc(
		fmt.Sprintf("%s/stats/hits", path),
		func(w http.ResponseWriter, r *http.Request) {
			hits := &goatcounter.HitsResponse{
				Hits: []goatcounter.Hit{
					{
						Count:  1,
						PathID: 1,
						Path:   "/",
						Event:  false,
						Title:  "Test",
						Max:    1,
						Stats: []goatcounter.Stat{
							{
								Day:   day,
								Daily: 1,
							},
						},
					},
				},
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(hits); err != nil {
				t.Fatalf("failed to encode hits response\n%v", err)
			}
		},
	)
	mux.HandleFunc(
		fmt.Sprintf("%s/stats/total", path),
		func(w http.ResponseWriter, r *http.Request) {
			total := &goatcounter.Total{
				Total:       1,
				TotalEvents: 1,
				TotalUTC:    60,
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(total); err != nil {
				t.Fatalf("failed to encode total response\n%v", err)
			}
		},
	)

	tmpl := template.Must(template.New("stats").Parse(prometheusStatistics))
	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, struct {
		Code     string
		Instance string
		Day      string
	}{
		Code:     code,
		Instance: instance,
		Day:      day,
	}); err != nil {
		t.Fatalf("failed to execute template: %v", err)
	}

	collector := NewStatisticsCollector(client, logger)
	if err := testutil.CollectAndCompare(
		collector,
		&buf,
	); err != nil {
		t.Errorf("unexpected error collecting result\n%v", err)
	}
}
