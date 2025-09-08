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
//! Provides REST API endpoints for querying logs and anomalies.

package api

import (
	"encoding/json"
	"net/http"

	"m2loganalyzer/internal/storage"
)

func QueryLogsHandler(db storage.ClickHouseStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO: implement query parsing
		rows, _ := db.db.Query("SELECT message, level, source FROM logs LIMIT 100")
		defer rows.Close()

		results := []map[string]string{}
		for rows.Next() {
			var msg, level, source string
			rows.Scan(&msg, &level, &source)
			results = append(results, map[string]string{
				"message": msg, "level": level, "source": source,
			})
		}

		json.NewEncoder(w).Encode(results)
	}
}
