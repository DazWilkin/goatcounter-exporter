package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
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
	code := os.Getenv(Code)
	if code == "" {
		log.Fatalf("Expected environment to contain  '%s' variable", Code)
	}
	token := os.Getenv(Token)
	if token == "" {
		log.Fatalf("Expected environment to contain '%s' variable", Token)
	}

	// For endpoint and metricsPath
	flag.Parse()

	if GitCommit == "" {
		log.Println("[main] GitCommit value unchanged: expected to be set during build")
	}
	if OSVersion == "" {
		log.Println("[main] OSVersion value unchanged: expected to be set during build")
	}

	client := goatcounter.NewClient(code, token)

	registry := prometheus.NewRegistry()
	registry.MustRegister(collector.NewExporterCollector(OSVersion, GoVersion, GitCommit, StartTime))
	registry.MustRegister(collector.NewPathsCollector(client))
	registry.MustRegister(collector.NewStatisticsCollector(client))

	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handleRoot))
	mux.Handle("/healthz", http.HandlerFunc(handleHealthz))
	mux.Handle(*metricsPath, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))

	log.Printf("[main] Server starting (%s)", *endpoint)
	log.Printf("[main] metrics served on: %s", *metricsPath)
	log.Fatal(http.ListenAndServe(*endpoint, mux))
}
