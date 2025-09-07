package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"m2loganalyzer/internal/config"
	"m2loganalyzer/internal/detector"
	"m2loganalyzer/internal/ingest"
	"m2loganalyzer/internal/pipeline"
)

func main() {
	log.Println("=== Starting M2 Log Analyzer AI v0.2 ===")

	// Load configuration
	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize rolling z-score detector
	det := detector.NewRollingZScoreDetector(cfg.Detector.WindowSize, cfg.Detector.Threshold)

	// Initialize pipeline processor with rules & detector
	proc := pipeline.NewProcessor(cfg.Pipeline.WorkerCount, cfg.Pipeline.QueueSize, &cfg.PipelineRules, det)
	proc.Start()

	// Initialize HTTP ingestor
	httpIngestor := ingest.NewHTTPIngestor(cfg, proc)

	// HTTP server with graceful shutdown
	server := &http.Server{
		Addr:    ":" + cfg.HTTP.Port,
		Handler: http.DefaultServeMux,
	}

	// Start HTTP server in goroutine
	go func() {
		log.Printf("HTTP ingestor listening on port %s", cfg.HTTP.Port)
		if err := httpIngestor.Start(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()

	// Listen for OS signals (Ctrl+C / SIGTERM)
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	log.Println("Shutdown signal received, cleaning up...")

	// Shutdown HTTP server gracefully
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("HTTP server shutdown error: %v", err)
	} else {
		log.Println("HTTP server stopped")
	}

	// Stop pipeline processor
	proc.Stop()

	log.Println("=== M2 Log Analyzer AI v0.2 stopped gracefully ===")
}
