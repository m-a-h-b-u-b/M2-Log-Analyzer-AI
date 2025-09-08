//! Module Name: main.go
//! --------------------------------
//! License : Apache 2.0
//! Author  : Md Mahbubur Rahman
//! URL     : https://m-a-h-b-u-b.github.io
//! GitHub  : https://github.com/m-a-h-b-u-b/M2-Log-Analyzer-AI
//!
//! Module Description:
//! Main entrypoint for v1.0 with multi-tenant support, pipelines,
//! ClickHouse/SQLite storage, API server, and Web UI.

package main

import (
	"log"
	"net/http"
	"path/filepath"

	"m2loganalyzer/internal/api"
	"m2loganalyzer/internal/config"
	"m2loganalyzer/internal/ingest"
	"m2loganalyzer/internal/multi_tenant"
	"m2loganalyzer/internal/pipeline"
	"m2loganalyzer/internal/storage"
	"m2loganalyzer/internal/util"
)

func main() {
	// Load global config
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Initialize Prometheus metrics
	util.InitMetrics()

	// Initialize tenant manager
	tenants := multi_tenant.NewTenantManager()

	// Example: load tenants from config (pseudo)
	for _, t := range cfg.Tenants {
		// Each tenant gets a pipeline and storage
		proc := pipeline.NewProcessor(t.Workers, t.QueueSize)
		go proc.Start()

		var store storage.Storage
		if t.UseClickHouse {
			ch, err := storage.NewClickHouse(t.ClickHouseDSN)
			if err != nil {
				log.Fatalf("tenant %s clickhouse init failed: %v", t.Name, err)
			}
			store = ch
		} else {
			sqlitePath := filepath.Join("data", t.Name+".db")
			sqlite, err := storage.NewSQLite(sqlitePath)
			if err != nil {
				log.Fatalf("tenant %s sqlite init failed: %v", t.Name, err)
			}
			store = sqlite
		}

		tenants.AddTenant(t.Name, proc, store)
	}

	// Initialize HTTP ingestion for all tenants
	httpIngest := ingest.NewHTTPIngestorMulti(tenants)
	go func() {
		log.Printf("HTTP ingestor listening on :%d", cfg.ServerPort)
		if err := http.ListenAndServe(
			":"+string(rune(cfg.ServerPort)), httpIngest.Router()); err != nil {
			log.Fatalf("HTTP server failed: %v", err)
		}
	}()

	// Start API server
	api.StartServerMulti(cfg, tenants)

	// Serve Web UI static files
	fs := http.FileServer(http.Dir("web/build"))
	http.Handle("/", fs)

	log.Println("M2-Log-Analyzer-AI v1.0 running!")
	select {} // block forever
}
