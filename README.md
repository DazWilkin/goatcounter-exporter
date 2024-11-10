# A Prometheus Exporter for [GoatCounter](https://goatcounter.com)

[![GitHub Actions](https://github.com/DazWilkin/goatcounter-exporter/actions/workflows/build.yml/badge.svg)](https://github.com/DazWilkin/goatcounter-exporter/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/DazWilkin/goatcounter-exporter.svg)](https://pkg.go.dev/github.com/DazWilkin/goatcounter-exporter)
[![Go Report Card](https://goreportcard.com/badge/github.com/DazWilkin/goatcounter-exporter)](https://goreportcard.com/report/github.com/DazWilkin/goatcounter-exporter)

## Run

The exporter requires your GoatCounter code and an API token provided by the environment. The exports accepts a flag for `endpoint` of the exporter, and metrics `path`:

```bash
export CODE="..."
export TOKEN="..."

HOST_PORT="8080"
CONT_PORT="8080"

podman run \
--interactive --tty --rm \
--name=goatcounter-exporter \
--env=CODE=${CODE} \
--env=TOKEN=${TOKEN} \
--publish=${HOST_PORT}:${CONT_PORT}/tcp \
ghcr.io/dazwilkin/goatcounter-exporter:25fbdf0c47c921d9ea24eecc00ec57fd50ce4218 \
--endpoint=:${CONT_PORT} \
--path=/metrics
```

## Prometheus

```bash
VERS="v2.46.0"

# Binds to host network to scrape GoatCounter Exporter
podman run \
--interactive --tty --rm \
--net=host \
--volume=${PWD}/prometheus.yml:/etc/prometheus/prometheus.yml \
--volume=${PWD}/rules.yml:/etc/alertmanager/rules.yml \
quay.io/prometheus/prometheus:${VERS} \
  --config.file=/etc/prometheus/prometheus.yml \
  --web.enable-lifecycle
```

## Metrics

All metrics are prefixed `goatcounter_exporter_`

|Name|Type|Description|
|----|----|-----------|
|`goatcounter_exporter_build_info`|Counter||
|`goatcounter_exporter_paths_total`|Gauge||
|`goatcounter_exporter_start_time`|Gauge||
|`goatcounter_exporter_stats_hits`|Gauge||
|`goatcounter_exporter_stats_total`|Gauge||

```
# HELP goatcounter_exporter_build_info A metric with a constant '1' value labeled by OS version, Go version, and the Git commit of the exporter
# TYPE goatcounter_exporter_build_info counter
goatcounter_exporter_build_info{git_commit="",go_version="",os_version=""}
# HELP goatcounter_exporter_paths_total List total of paths
# TYPE goatcounter_exporter_paths_total gauge
goatcounter_exporter_paths_total
# HELP goatcounter_exporter_start_time Exporter start time in Unix epoch seconds
# TYPE goatcounter_exporter_start_time gauge
goatcounter_exporter_start_time
# HELP goatcounter_exporter_stats_hits pageview and visitor stats
# TYPE goatcounter_exporter_stats_hits gauge
goatcounter_exporter_stats_hits{day="",path=""}
# HELP goatcounter_exporter_stats_total List total pageview counts
# TYPE goatcounter_exporter_stats_total gauge
goatcounter_exporter_stats_total
```

## [Sigstore](https://www.sigstore.dev/)

`goatcounter-exporter` container images are being signed by [Sigstore](https://www.sigstore.dev/) and may be verified:

```bash
cosign verify \
--key=./cosign.pub \
ghcr.io/dazwilkin/goatcounter-exporter:25fbdf0c47c921d9ea24eecc00ec57fd50ce4218
```

> **NOTE** `cosign.pub` may be downloaded [here](https://github.com/DazWilkin/goatcounter-exporter/blob/master/cosign.pub)

To install `cosign`, e.g.:
```bash
go install github.com/sigstore/cosign/cmd/cosign@latest
```
## Similar Exporters

+ [Prometheus Exporter for Azure](https://github.com/DazWilkin/azure-exporter)
+ [Prometheus Exporter for Fly.io](https://github.com/DazWilkin/fly-exporter)
+ [Prometheus Exporter for GCP](https://github.com/DazWilkin/gcp-exporter)
+ [Prometheus Exporter for Koyeb](https://github.com/DazWilkin/koyeb-exporter)
+ [Prometheus Exporter for Linode](https://github.com/DazWilkin/linode-exporter)
+ [Prometheus Exporter for Porkbun](https://github.com/DazWilkin/porkbun-exporter)
+ [Prometheus Exporter for Vultr](https://github.com/DazWilkin/vultr-exporter)


<hr/>
<br/>
<a href="https://www.buymeacoffee.com/dazwilkin" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy Me A Coffee" height="41" width="174"></a>