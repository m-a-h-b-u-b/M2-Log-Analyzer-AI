package ingest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"m2loganalyzer/internal/config"
	"m2loganalyzer/internal/pipeline"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// HTTPIngestor handles HTTP log ingestion
type HTTPIngestor struct {
	cfg  *config.Config
	proc *pipeline.Processor
}

// NewHTTPIngestor creates a new HTTP ingestor
func NewHTTPIngestor(cfg *config.Config, proc *pipeline.Processor) *HTTPIngestor {
	return &HTTPIngestor{
		cfg:  cfg,
		proc: proc,
	}
}

// Start runs the HTTP server for ingestion
func (h *HTTPIngestor) Start() error {
	http.HandleFunc("/ingest", h.handleIngest)

	addr := fmt.Sprintf(":%s", h.cfg.HTTP.Port)
	log.Printf("HTTP ingestor listening on %s", addr)
	return http.ListenAndServe(addr, nil)
}

// handleIngest receives logs via POST and submits them to the pipeline
func (h *HTTPIngestor) handleIngest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	logLine := string(body)
	h.proc.Submit(logLine)

	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte("Log accepted"))
}

func (h *HTTPIngestor) Start() error {
	// Existing ingest handler
	http.HandleFunc("/ingest", h.handleIngest)

	// Add /metrics endpoint for Prometheus
	http.Handle("/metrics", promhttp.Handler())

	addr := fmt.Sprintf(":%s", h.cfg.HTTP.Port)
	log.Printf("HTTP ingestor listening on %s", addr)
	return http.ListenAndServe(addr, nil)
}