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
//! API server exposing health, metrics, and model retraining endpoints.

package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"m2loganalyzer/internal/pipeline"
)

// retrainResponse defines the JSON response for retrain requests
type retrainResponse struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Time    string `json:"time"`
}

func StartServer(cfg interface{}, proc *pipeline.Processor) {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	http.HandleFunc("/retrain", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
			return
		}

		// Respond immediately, run retraining in background
		go func() {
			log.Println("[Retrain] Starting retraining process...")
			err := proc.Retrain() // <-- you need to implement this inside pipeline.Processor
			if err != nil {
				log.Printf("[Retrain] Failed: %v\n", err)
			} else {
				log.Println("[Retrain] Completed successfully.")
			}
		}()

		resp := retrainResponse{
			Status: "accepted",
			Time:   time.Now().Format(time.RFC3339),
			Message: "Retraining has been triggered. " +
				"Check logs for progress.",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	})

	log.Println("API server running on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
