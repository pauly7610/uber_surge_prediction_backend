package utils

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	requestLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "request_latency_seconds",
			Help: "Request latency in seconds.",
		},
		[]string{"endpoint"},
	)
	errorRates = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "error_rates",
			Help: "Error rates by endpoint.",
		},
		[]string{"endpoint"},
	)
	kafkaLag = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "kafka_lag",
			Help: "Kafka consumer lag.",
		},
		[]string{"topic"},
	)
)

func init() {
	prometheus.MustRegister(requestLatency, errorRates, kafkaLag)
} 