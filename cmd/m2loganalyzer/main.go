package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"m2loganalyzer/internal/config"
	"m2loganalyzer/internal/ingest"
	"m2loganalyzer/internal/pipeline"
)

func main() {
	// Load config
	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Start pipeline processor
	proc := pipeline.NewProcessor(cfg.Pipeline.WorkerCount, cfg.Pipeline.QueueSize)
	proc.Start()

	// Start HTTP ingestor
	httpIngestor := ingest.NewHTTPIngestor(cfg, proc)
	go func() {
		if err := httpIngestor.Start(); err != nil {
			log.Fatalf("HTTP ingestor error: %v", err)
		}
	}()

	// Wait for interrupt signal
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	<-sigCh

	log.Println("Shutting down...")
	proc.Stop()
}
