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
//! Worker pool processor that receives log events, processes them,
//! and forwards to detectors, storage, and alerts.

package pipeline

import (
	"log"
	"sync"

	"m2loganalyzer/internal/util"
)

type LogEvent struct {
	Message string `json:"message"`
	Level   string `json:"level"`
	Source  string `json:"source"`
}

type Processor struct {
	workers int
	queue   chan LogEvent
	wg      sync.WaitGroup
}

func NewProcessor(workers, queueSize int) *Processor {
	return &Processor{
		workers: workers,
		queue:   make(chan LogEvent, queueSize),
	}
}

func (p *Processor) Start() {
	for i := 0; i < p.workers; i++ {
		p.wg.Add(1)
		go func(id int) {
			defer p.wg.Done()
			for event := range p.queue {
				log.Printf("[worker %d] processing log: %+v", id, event)
				util.IncLogsProcessed()
			}
		}(i)
	}
}

func (p *Processor) Submit(event LogEvent) {
	select {
	case p.queue <- event:
	default:
		util.IncLogsDropped()
	}
}

func (p *Processor) Stop() {
	close(p.queue)
	p.wg.Wait()
}
