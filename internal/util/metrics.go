package util

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// Total logs received by HTTP ingestor
	LogsReceived = promauto.NewCounter(prometheus.CounterOpts{
		Name: "m2_logs_received_total",
		Help: "Total number of logs received by HTTP ingestor",
	})

	// Total logs successfully processed by pipeline
	LogsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "m2_logs_processed_total",
		Help: "Total number of logs processed by pipeline",
	})

	// Total logs dropped due to full queue
	LogsDropped = promauto.NewCounter(prometheus.CounterOpts{
		Name: "m2_logs_dropped_total",
		Help: "Total number of logs dropped due to full queue",
	})
)
