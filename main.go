package main

import (
	"flag"
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/DazWilkin/goatcounter-exporter/collector"
	"github.com/DazWilkin/goatcounter-exporter/goatcounter"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	Code  string = "CODE"
	Token string = "TOKEN"
)
const (
	sRoot string = `
<h2>A Prometheus Exporter for <a href="https://goatcounter.com">GoatCounter</a></h2>
<ul>
	<li><a href="{{ .Metrics }}">metrics</a></li>
	<li><a href="/healthz">healthz</a></li>
</ul>`
)

var (
	// GitCommit is the git commit value and is expected to be set during build
	GitCommit string
	// GoVersion is the Golang runtime version
	GoVersion = runtime.Version()
	// OSVersion is the OS version (uname --kernel-release) and is expected to be set during build
	OSVersion string
	// StartTime is the start time of the exporter represented as a UNIX epoch
	StartTime = time.Now().Unix()
)
var (
	endpoint    = flag.String("endpoint", ":8080", "The endpoint of the Exporter HTTP server")
	instance    = flag.String("instance", "goatcounter.com", "The endpoint of the GoatCounter API")
	metricsPath = flag.String("path", "/metrics", "The path on which Prometheus metrics will be served")
)
var (
	tRoot = template.Must(template.New("root").Parse(sRoot))
)

func handleHealthz(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}
func handleRoot(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	fmt.Fprint(w)
	tRoot.Execute(w, struct {
		Metrics string
	}{
		Metrics: *metricsPath,
	})
}
func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// CODE represents
	// Either a user-defined host on goatcounter.com
	// Or, in combination with the instance flag, a host on a user-defined domain
	code := os.Getenv(Code)
	if code == "" {
		logger.Info("Expected environment to contain 'CODE' variable",
			"variable", Code,
		)
	}
	token := os.Getenv(Token)
	if token == "" {
		logger.Info("Expected environment to contain 'TOKEN' variable",
			"variable", Token,
		)
	}

	// For endpoint, instance and metricsPath
	flag.Parse()

	if GitCommit == "" {
		logger.Info("GitCommit value unchanged: expected to be set during build")
	}
	if OSVersion == "" {
		logger.Info("OSVersion value unchanged: expected to be set during build")
	}

	goatcounterEndpoint := fmt.Sprintf("%s.%s", code, *instance)
	client := goatcounter.NewClient(goatcounterEndpoint, token)

	registry := prometheus.NewRegistry()
	registry.MustRegister(collector.NewExporterCollector(OSVersion, GoVersion, GitCommit, StartTime))
	registry.MustRegister(collector.NewPathsCollector(client, logger))
	registry.MustRegister(collector.NewStatisticsCollector(client, logger))

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handleRoot))
	mux.Handle("/healthz", http.HandlerFunc(handleHealthz))
	mux.Handle(*metricsPath, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))

	logger.Info("Server starting",
		"endpoint", *endpoint,
	)
	logger.Info("metrics path",
		"path", *metricsPath,
	)
	logger.Error("Server failed", http.ListenAndServe(*endpoint, mux))
}
