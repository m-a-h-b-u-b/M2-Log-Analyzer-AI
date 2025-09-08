//! Module Name: http_multi.go
//! --------------------------------
//! License : Apache 2.0
//! Author  : Md Mahbubur Rahman
//! URL     : https://m-a-h-b-u-b.github.io
//! GitHub  : https://github.com/m-a-h-b-u-b/M2-Log-Analyzer-AI
//!
//! Module Description:
//! HTTP ingestor for multi-tenant log ingestion.

package ingest

import (
	"encoding/json"
	"net/http"

	"m2loganalyzer/internal/multi_tenant"
	"m2loganalyzer/internal/pipeline"
	"m2loganalyzer/internal/util"
)

type HTTPIngestorMulti struct {
	manager *multi_tenant.TenantManager
}

func NewHTTPIngestorMulti(manager *multi_tenant.TenantManager) *HTTPIngestorMulti {
	return &HTTPIngestorMulti{manager: manager}
}

func (h *HTTPIngestorMulti) Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/ingest", h.handleIngest)
	mux.Handle("/metrics", util.PrometheusHandler())
	return mux
}

type MultiLogEvent struct {
	Tenant  string         `json:"tenant"`
	LogEvent pipeline.LogEvent `json:"log"`
}

func (h *HTTPIngestorMulti) handleIngest(w http.ResponseWriter, r *http.Request) {
	var event MultiLogEvent
	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "invalid log format", http.StatusBadRequest)
		return
	}

	tenant, ok := h.manager.GetTenant(event.Tenant)
	if !ok {
		http.Error(w, "tenant not found", http.StatusNotFound)
		return
	}

	tenant.Pipeline.Submit(event.LogEvent)
	util.IncLogsReceived()
	w.WriteHeader(http.StatusAccepted)
}
