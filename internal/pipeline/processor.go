package pipeline

import (
	"context"
	"log"
	"sync"
	"m2loganalyzer/internal/util"
	"m2loganalyzer/internal/detector"
	"m2loganalyzer/internal/config"
	"m2loganalyzer/internal/util"
)

// Processor handles log events with a worker pool
type Processor struct {
	jobs       chan string
	workerCount int
	wg         sync.WaitGroup
	ctx        context.Context
	cancel     context.CancelFunc
	rules   *config.PipelineRules
	detector *detector.RollingZScoreDetector
}

// NewProcessor creates a new processor
func NewProcessor(workerCount, queueSize int) *Processor {
	ctx, cancel := context.WithCancel(context.Background())
	return &Processor{
		jobs:        make(chan string, queueSize),
		workerCount: workerCount,
		ctx:         ctx,
		cancel:      cancel,
	}
}

// Start launches worker goroutines
func (p *Processor) Start() {
	for i := 0; i < p.workerCount; i++ {
		p.wg.Add(1)
		go p.worker(i)
	}
	log.Printf("Started %d pipeline workers", p.workerCount)
}

// worker processes log events
func (p *Processor) worker(id int) {
	defer p.wg.Done()
	for {
		select {
		case <-p.ctx.Done():
			log.Printf("Worker %d shutting down...", id)
			return
		case logLine := <-p.jobs:
			p.processLog(logLine)
		}
	}
}

// processLog is a placeholder for actual processing logic
func (p *Processor) processLog(logLine string) {
	// TODO: parsing, enrichment, detector, alerting
	log.Printf("Processing log: %s", logLine)
}

// Submit adds a log line to the queue
func (p *Processor) Submit(logLine string) {
	select {
	case p.jobs <- logLine:
	default:
		log.Println("Dropping log (queue full)")
	}
}

// Stop gracefully stops all workers
func (p *Processor) Stop() {
	p.cancel()
	p.wg.Wait()
	log.Println("All workers stopped")
}

func (p *Processor) processLog(logLine string) {
	// Increment processed counter
	util.LogsProcessed.Inc()

	// TODO: parsing, enrichment, detector, alerting
	log.Printf("Processing log: %s", logLine)
}

func (p *Processor) Submit(logLine string) {
	select {
	case p.jobs <- logLine:
		util.LogsReceived.Inc() // Count received logs
	default:
		util.LogsDropped.Inc()  // Count dropped logs
		log.Println("Dropping log (queue full)")
	}
}

func NewProcessor(workerCount, queueSize int, rules *config.PipelineRules, det *detector.RollingZScoreDetector) *Processor {
	ctx, cancel := context.WithCancel(context.Background())
	return &Processor{
		jobs:        make(chan string, queueSize),
		workerCount: workerCount,
		ctx:         ctx,
		cancel:      cancel,
		rules:       rules,
		detector:    det,
	}
}

func (p *Processor) processLog(logLine string) {
	util.LogsProcessed.Inc()

	// Convert log to float for demo purposes (e.g., metric value)
	val := float64(len(logLine)) // Example: log length as value

	// Apply rolling z-score
	if p.detector.AddValue(val) {
		log.Printf("Anomaly detected: %s", logLine)
		if p.rules.SlackWebhook != "" {
			go alertSlack(p.rules.SlackWebhook, logLine)
		}
		if p.rules.Webhook != "" {
			go alertWebhook(p.rules.Webhook, logLine)
		}
	}
}