package main

import (
	"log"
	"net/http"
	_ "net/http"

	"github.com/petrolax/yaml-converter/convert"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	_ "github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())

	metrics := convert.YamlToOpenMetrics("assets/test.yaml")
	for _, metric := range metrics {
		prometheus.MustRegister(metric)
	}

	log.Fatal(http.ListenAndServe(":8080", nil))
}
