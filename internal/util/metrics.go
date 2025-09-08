//! M2 Log Analyzer AI
//! --------------------------------
//! License : Dual License
//!           - Apache 2.0 for open-source / personal use
//!           - Commercial license required for closed-source use
//! Author  : Md Mahbubur Rahman
//! URL     : https://m-a-h-b-u-b.github.io
//! GitHub  : https://github.com/m-a-h-b-u-b/M2-Log-Analyzer-AI
//!
//! Module Description:
//! Prometheus metrics collector for logs received, processed, and dropped.

package util

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	logsReceived = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "logs_received_total",
		Help: "Total number of logs received",
	})
	logsProcessed = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "logs_processed_total",
		Help: "Total number of logs processed",
	})
	logsDropped = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "logs_dropped_total",
		Help: "Total number of logs dropped due to queue overflow",
	})
)

func InitMetrics() {
	prometheus.MustRegister(logsReceived, logsProcessed, logsDropped)
}

func IncLogsReceived() { logsReceived.Inc() }
func IncLogsProcessed() { logsProcessed.Inc() }
func IncLogsDropped()  { logsDropped.Inc() }

func PrometheusHandler() http.Handler {
	return promhttp.Handler()
}
