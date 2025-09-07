# M2 Log Analyzer AI — Architecture Overview

## Components
- **Ingestors**: Collect logs (HTTP, syslog, file, Kafka/NATS)
- **Pipeline**: Parse, enrich, and route logs
- **Detectors**: Apply anomaly detection algorithms
- **Alerts**: Notify via Slack, webhook, email, or Prometheus
- **Storage**: Optional persistence (SQLite, ClickHouse, memory)
- **API**: Expose endpoints for ingestion, health, metrics, alerts

## Flow
Log Sources ➝ Ingestors ➝ Router ➝ Processor ➝ Detector ➝ Alerts/Storage
