//! Module Name: server.go
//! --------------------------------
//! License : Apache 2.0
//! Author  : Md Mahbubur Rahman
//! URL     : https://m-a-h-b-u-b.github.io
//! GitHub  : https://github.com/m-a-h-b-u-b/M2-Log-Analyzer-AI
//!
//! Module Description:
//! API server exposing health, metrics, and model retraining endpoints.

package api

import (
	"fmt"
	"log"
	"net/http"

	"m2loganalyzer/internal/pipeline"
)

func StartServer(cfg interface{}, proc *pipeline.Processor) {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	http.HandleFunc("/retrain", func(w http.ResponseWriter, r *http.Request) {
		// TODO: implement retraining logic
		fmt.Fprintln(w, "retrain triggered")
	})

	log.Println("API server running on :8080")
	http.ListenAndServe(":8080", nil)
}
