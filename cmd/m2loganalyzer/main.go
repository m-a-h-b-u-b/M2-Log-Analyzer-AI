package main

import (
	"fmt"
	"log"

	"m2loganalyzer/internal/config"
)

func main() {
	cfg, err := config.Load("configs/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	fmt.Printf("HTTP Port: %s\n", cfg.HTTP.Port)
	fmt.Printf("Pipeline Workers: %d\n", cfg.Pipeline.WorkerCount)
	fmt.Printf("Detector Type: %s\n", cfg.Detector.Type)
	fmt.Printf("Slack Webhook: %s\n", cfg.Alerts.SlackWebhook)
}
