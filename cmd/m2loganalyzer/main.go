//! Module Name: main.go
//! --------------------------------
//! License : Apache 2.0
//! Author  : Md Mahbubur Rahman
//! URL     : https://m-a-h-b-u-b.github.io
//! GitHub  : https://github.com/m-a-h-b-u-b/M2-Log-Analyzer-AI
//!
//! Module Description:
//! Main entrypoint that wires configuration, ingestion, pipeline,
//! anomaly detectors, storage backends, and alerting.

package main

import (
	"log"
	"net/http"

	"m2loganalyzer/internal/api"
	"m2loganalyzer/internal/config"
	"m2loganalyzer/internal/ingest"
	"m2loganalyzer/internal/pipeline"
	"m2loganalyzer/internal/util"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Init metrics
	util.InitMetrics()

	// Init pipeline
	proc := pipeline.NewProcessor(cfg.Workers, cfg.QueueSize)
	go proc.Start()

	// Init HTTP ingestor
	httpIngest := ingest.NewHTTPIngestor(proc)
	go func() {
		log.Printf("HTTP ingestor listening on :%d", cfg.ServerPort)
		if err := http.ListenAndServe(
			":"+string(rune(cfg.ServerPort)), httpIngest.Router()); err != nil {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()

	// Init API server
	api.StartServer(cfg, proc)
}
