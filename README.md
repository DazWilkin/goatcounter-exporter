# A Prometheus Exporter for [GoatCounter](https://goatcounter.com)

[![build](https://github.com/DazWilkin/goatcounter-exporter/actions/workflows/build.yml/badge.svg)](https://github.com/DazWilkin/goatcounter-exporter/actions/workflows/build.yml)

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
ghcr.io/dazwilkin/goatcounter-exporter:2f87ab6a277958ea036d0e9e078bd577a4f82f6f \
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

## [Sigstore](https://www.sigstore.dev/)

`goatcounter-exporter` container images are being signed by [Sigstore](https://www.sigstore.dev/) and may be verified:

```bash
cosign verify \
--key=./cosign.pub \
ghcr.io/dazwilkin/goatcounter-exporter:2f87ab6a277958ea036d0e9e078bd577a4f82f6f
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