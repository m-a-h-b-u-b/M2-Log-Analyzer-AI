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
//! HTTP ingestor that accepts log events via POST and feeds them into the pipeline.

package ingest

import (
	"encoding/json"
	"net/http"

	"m2loganalyzer/internal/pipeline"
	"m2loganalyzer/internal/util"
)

type HTTPIngestor struct {
	proc *pipeline.Processor
}

func NewHTTPIngestor(proc *pipeline.Processor) *HTTPIngestor {
	return &HTTPIngestor{proc: proc}
}

func (h *HTTPIngestor) Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ingest", h.handleIngest)
	mux.Handle("/metrics", util.PrometheusHandler())
	return mux
}

func (h *HTTPIngestor) handleIngest(w http.ResponseWriter, r *http.Request) {
	var event pipeline.LogEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "invalid log format", http.StatusBadRequest)
		return
	}
	h.proc.Submit(event)
	util.IncLogsReceived()
	w.WriteHeader(http.StatusAccepted)
}
