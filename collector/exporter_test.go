package collector

import (
	"strings"
	"testing"

	"github.com/prometheus/client_golang/prometheus/testutil"
)

const (
	// prometheusExporter is a string; conveniently start_time is always rendered as 1 for testing
	prometheusExporter string = `
# HELP goatcounter_exporter_build_info A metric with a constant '1' value labeled by OS version, Go version, and the Git commit of the exporter
# TYPE goatcounter_exporter_build_info counter
goatcounter_exporter_build_info{git_commit="commit",go_version="go",os_version="os"} 1
# HELP goatcounter_exporter_start_time Exporter start time in Unix epoch seconds
# TYPE goatcounter_exporter_start_time gauge
goatcounter_exporter_start_time 1
`
)

func TestExporterCollector(t *testing.T) {
	collector := NewExporterCollector("os", "go", "commit", 1)
	if err := testutil.CollectAndCompare(
		collector,
		strings.NewReader(prometheusExporter),
	); err != nil {
		t.Errorf("unexpected error collecting result\n%v", err)
	}
}
