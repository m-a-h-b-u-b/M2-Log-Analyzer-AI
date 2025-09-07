package main

import (
	"fmt"
	"log"

	"m2loganalyzer/internal/config"
	"m2loganalyzer/internal/pipeline"
)

func main() {
	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Create processor with config values
	proc := pipeline.NewProcessor(cfg.Pipeline.WorkerCount, cfg.Pipeline.QueueSize)
	proc.Start()

	// Submit some test logs
	for i := 0; i < 10; i++ {
		proc.Submit(fmt.Sprintf("Test log #%d", i+1))
	}

	// Wait for user input to stop
	fmt.Println("Press Enter to stop...")
	fmt.Scanln()
	proc.Stop()
}
