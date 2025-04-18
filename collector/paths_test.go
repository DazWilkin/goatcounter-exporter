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

	"github.com/DazWilkin/goatcounter-exporter/goatcounter"
	"github.com/prometheus/client_golang/prometheus/testutil"
)

const (
	prometheusPaths string = `
# HELP goatcounter_exporter_paths_total List total of paths
# TYPE goatcounter_exporter_paths_total gauge
goatcounter_exporter_paths_total{code="{{ .Code }}",instance="{{ .Instance }}"} 1
`
)

// TestPathsCollector tests PathsCollector.Collect
// Collect covers 1 API (/paths/list) represented here by Handler closures
func TestPathsCollector(t *testing.T) {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	defer server.Close()

	code := "test"
	instance := server.URL

	client := goatcounter.NewTestClient(code, instance, logger)
	path := fmt.Sprintf("/%s", goatcounter.Version)

	mux.HandleFunc(
		fmt.Sprintf("%s/paths/", path),
		func(w http.ResponseWriter, r *http.Request) {
			paths := &goatcounter.PathsResponse{
				Paths: []goatcounter.Path{
					{
						Event: false,
						Path:  "/",
						ID:    0,
						Title: "test",
					},
				},
				More: false,
			}
			w.Header().Set("Content-Type", "application/json")
			if err := json.NewEncoder(w).Encode(paths); err != nil {
				t.Fatalf("failed to encode hits response\n%v", err)
			}
		},
	)

	tmpl := template.Must(template.New("paths").Parse(prometheusPaths))
	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, struct {
		Code     string
		Instance string
	}{
		Code:     code,
		Instance: instance,
	}); err != nil {
		t.Fatalf("failed to execute template: %v", err)
	}

	collector := NewPathsCollector(client, logger)
	if err := testutil.CollectAndCompare(
		collector,
		&buf,
	); err != nil {
		t.Errorf("unexpected error collecting result\n%v", err)
	}
}
