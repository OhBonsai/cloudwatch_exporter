package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func metricsHandler(w http.ResponseWriter, req *http.Request) {
	tagsData, cloudwatchData := scrapeAwsData(config)

	var metrics []*PrometheusMetric

	metrics = append(metrics, migrateCloudwatchToPrometheus(cloudwatchData)...)
	metrics = append(metrics, migrateTagsToPrometheus(tagsData)...)

	registry := prometheus.NewRegistry()
	registry.MustRegister(NewPrometheusCollector(metrics))

	if err := registry.Register(cloudwatchAPICounter); err != nil {
		log.Fatal("Could not publish cloudwatch api metric")
	}

	handler := promhttp.HandlerFor(registry, promhttp.HandlerOpts{
		DisableCompression: false,
	})

	handler.ServeHTTP(w, req)
}
