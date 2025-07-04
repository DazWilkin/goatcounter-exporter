# A Prometheus Exporter for [GoatCounter](https://goatcounter.com)

[![GitHub Actions](https://github.com/DazWilkin/goatcounter-exporter/actions/workflows/build.yml/badge.svg)](https://github.com/DazWilkin/goatcounter-exporter/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/DazWilkin/goatcounter-exporter.svg)](https://pkg.go.dev/github.com/DazWilkin/goatcounter-exporter)
[![Go Report Card](https://goreportcard.com/badge/github.com/DazWilkin/goatcounter-exporter)](https://goreportcard.com/report/github.com/DazWilkin/goatcounter-exporter)

## Run

The exporter requires your GoatCounter code and an API token provided by the environment. The exports accepts a flag for `endpoint` of the exporter, an optional `instance` endpoint of the GoatCounter API, and a `metricsPath`:

```bash
CODE="..."
TOKEN="..."

HOST_PORT="8080"
CONT_PORT="8080"

podman run \
--interactive --tty --rm \
--name=goatcounter-exporter \
--env=CODE=${CODE} \
--env=TOKEN=${TOKEN} \
--publish=${HOST_PORT}:${CONT_PORT}/tcp \
ghcr.io/dazwilkin/goatcounter-exporter:664658bc4b8b2801f4aa735caf15f747e63a0e55 \
--endpoint=:${CONT_PORT} \
--instance="goatcounter.com" \
--path="/metrics"
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
ghcr.io/dazwilkin/goatcounter-exporter:664658bc4b8b2801f4aa735caf15f747e63a0e55
```

> **NOTE** `cosign.pub` may be downloaded [here](https://github.com/DazWilkin/goatcounter-exporter/blob/master/cosign.pub)

To install `cosign`, e.g.:
```bash
go install github.com/sigstore/cosign/cmd/cosign@latest
```
## Similar Exporters

+ [Prometheus Exporter for Azure](https://github.com/DazWilkin/azure-exporter)
+ [Prometheus Exporter for crt.sh](https://github.com/DazWilkin/crtsh-exporter)
+ [Prometheus Exporter for Fly.io](https://github.com/DazWilkin/fly-exporter)
+ [Prometheus Exporter for GoatCounter](https://github.com/DazWilkin/goatcounter-exporter)
+ [Prometheus Exporter for Google Cloud](https://github.com/DazWilkin/gcp-exporter)
+ [Prometheus Exporter for Koyeb](https://github.com/DazWilkin/koyeb-exporter)
+ [Prometheus Exporter for Linode](https://github.com/DazWilkin/linode-exporter)
+ [Prometheus Exporter for PorkBun](https://github.com/DazWilkin/porkbun-exporter)
+ [Prometheus Exporter for updown.io](https://github.com/DazWilkin/updown-exporter)
+ [Prometheus Exporter for Vultr](https://github.com/DazWilkin/vultr-exporter)

<hr/>
<br/>
<a href="https://www.buymeacoffee.com/dazwilkin" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/default-orange.png" alt="Buy Me A Coffee" height="41" width="174"></a>
