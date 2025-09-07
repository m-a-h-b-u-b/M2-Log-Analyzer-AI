# M2 Log Analyzer AI

Lightweight AI-powered log ingestion and anomaly detection system written in Go.  
A modern, minimal alternative to ELK, focusing on concurrency, stream processing, and embedded ML.

---

## 🚀 Features (planned)
- Real-time log ingestion (HTTP, syslog, file, Kafka/NATS)
- Concurrency-first design (worker pools, Go channels)
- Stream parsing & enrichment
- Anomaly detection (rolling stats, Isolation Forest, Autoencoder)
- Alerts (Slack, webhook, email, Prometheus)
- Optional storage (SQLite, ClickHouse, in-memory)
- Lightweight deployment (Docker + Kubernetes)

---

## 📂 Project Structure
```text
cmd/            # Entrypoint (main.go)
internal/       # Core logic (ingestion, pipeline, detectors, alerts, storage)
pkg/            # Public reusable packages
configs/        # Example configuration files
deploy/         # Docker + Kubernetes manifests
examples/       # Sample logs & usage examples
docs/           # Documentation & architecture notes
tests/          # Integration and load tests
