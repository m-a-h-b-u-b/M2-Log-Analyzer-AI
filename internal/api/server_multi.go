//! Module Name: server_multi.go
//! --------------------------------
//! License : Apache 2.0
//! Author  : Md Mahbubur Rahman
//! URL     : https://m-a-h-b-u-b.github.io
//! GitHub  : https://github.com/m-a-h-b-u-b/M2-Log-Analyzer-AI
//!
//! Module Description:
//! Multi-tenant API server exposing health, metrics, retraining, and query endpoints.

package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"m2loganalyzer/internal/multi_tenant"
	"m2loganalyzer/internal/storage"
)

func StartServerMulti(cfg interface{}, manager *multi_tenant.TenantManager) {
	mux := http.NewServeMux()

	// Health endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	// Model retraining endpoint per tenant
	mux.HandleFunc("/retrain", func(w http.ResponseWriter, r *http.Request) {
		tenantID := r.URL.Query().Get("tenant")
		tenant, ok := manager.GetTenant(tenantID)
		if !ok {
			http.Error(w, "tenant not found", http.StatusNotFound)
			return
		}
		// TODO: implement retraining logic per tenant
		fmt.Fprintf(w, "retrain triggered for tenant %s\n", tenant.Name)
	})

	// Query logs endpoint per tenant
	mux.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		tenantID := r.URL.Query().Get("tenant")
		tenant, ok := manager.GetTenant(tenantID)
		if !ok {
			http.Error(w, "tenant not found", http.StatusNotFound)
			return
		}
		rows, _ := tenant.Store.(*storage.ClickHouseStorage).DB().Query("SELECT message, level, source FROM logs LIMIT 100")
		defer rows.Close()

		results := []map[string]string{}
		for rows.Next() {
			var msg, level, source string
			rows.Scan(&msg, &level, &source)
			results = append(results, map[string]string{"message": msg, "level": level, "source": source})
		}
		json.NewEncoder(w).Encode(results)
	})

	// Metrics endpoint
	mux.Handle("/metrics", util.PrometheusHandler())

	log.Println("Multi-tenant API server running on :8080")
	http.ListenAndServe(":8080", mux)
}
