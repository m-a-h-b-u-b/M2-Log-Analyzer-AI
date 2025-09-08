//! Module Name: tenant.go
//! --------------------------------
//! License : Apache 2.0
//! Author  : Md Mahbubur Rahman
//! URL     : https://m-a-h-b-u-b.github.io
//! GitHub  : https://github.com/m-a-h-b-u-b/M2-Log-Analyzer-AI
//!
//! Module Description:
//! Defines a single tenant with its pipeline and storage backend.

package multi_tenant

import (
	"m2loganalyzer/internal/pipeline"
	"m2loganalyzer/internal/storage"
)

type Tenant struct {
	Name     string
	Pipeline *pipeline.Processor
	Store    storage.Storage
}
