
# M2-Log-Analyzer-AI

[![Go](https://img.shields.io/badge/Go-1.21-blue?style=flat-square)](https://golang.org/)
[![Kafka](https://img.shields.io/badge/Kafka-Event%20Streaming-orange?style=flat-square)](https://kafka.apache.org/)
[![Kubernetes](https://img.shields.io/badge/Kubernetes-Helm-blue?style=flat-square)](https://kubernetes.io/)
[![License](https://img.shields.io/badge/License-Apache--2.0-green?style=flat-square)](https://opensource.org/licenses/Apache-2.0)


## Overview

**M2-Log-Analyzer-AI** is a **lightweight, AI-powered log analysis system** built in **Go**, designed for **real-time ingestion, stream processing, anomaly detection, alerting, and analytics**.  

It supports **multi-tenant deployments, scalable storage backends, and enterprise-ready dashboards**, making it a minimal yet powerful alternative to ELK.

---

## Contact

* **Author**: Md Mahbubur Rahman
* **GitHub**: [https://github.com/m-a-h-b-u-b](https://github.com/m-a-h-b-u-b)
* **Website/Portfolio**: [https://m-a-h-b-u-b.github.io](https://m-a-h-b-u-b.github.io)

---

## Features

- **Real-time ingestion:** HTTP, file tailing, syslog, Kafka/NATS  
- **Concurrent processing:** Worker pool handles thousands of log events/sec  
- **AI-powered detection:** Z-score and Isolation Forest models  
- **Alerts:** Slack, webhook, Prometheus push  
- **Metrics:** Logs received, processed, dropped  
- **Storage:** SQLite (lightweight) or ClickHouse (analytics)  
- **Multi-tenant:** Separate queues, pipelines, and databases per tenant  
- **Web UI:** Dashboard and query interface  
- **Config-driven:** Everything configurable via `config.yaml`  
- **Deployment-ready:** Docker + Helm for Kubernetes  

---

## Folder Structure

```
M2-Log-Analyzer-AI/
├── cmd/                        # Main entrypoint
│   └── m2loganalyzer/
├── internal/
│   ├── ingest/                 # Log ingestion
│   ├── pipeline/               # Worker pool & processing
│   ├── detector/               # AI anomaly detection
│   ├── storage/                # SQLite / ClickHouse / memory
│   ├── alert/                  # Slack / webhook / email / Prometheus
│   ├── api/                    # HTTP API endpoints
│   ├── config/                 # YAML config loader
│   ├── util/                   # Metrics, logging, tracing
│   └── multi_tenant/           # Multi-tenant management
├── configs/
│   └── config.yaml
├── deploy/
│   ├── docker/
│   │   └── Dockerfile
│   └── k8s/
│       ├── deployment.yaml
│       └── service.yaml
├── helm/                       # Helm chart for v1.0
│   ├── Chart.yaml
│   ├── values.yaml
│   └── templates/
├── web/                        # Web UI
│   ├── src/
│   └── public/
├── examples/
├── docs/
├── tests/
├── go.mod
├── go.sum
└── README.md
```

---

## Installation

### Prerequisites
- Go >= 1.21  
- SQLite3 / ClickHouse (optional)  
- Docker & Kubernetes (optional for deployment)  

### Build
```bash
git clone https://github.com/m-a-h-b-u-b/M2-Log-Analyzer-AI.git
cd M2-Log-Analyzer-AI
go mod tidy
go build -o m2loganalyzer ./cmd/m2loganalyzer
```

### Run
```bash
./m2loganalyzer
```

Metrics available at: `http://localhost:8080/metrics`  

---

## Usage Examples

### HTTP Ingestion
```bash
curl -X POST http://localhost:8080/ingest -H "Content-Type: application/json" -d '{"tenant":"tenant1","log":{"message":"User login failed","level":"error","source":"auth-service"}}'
```

### Query Logs
```bash
curl http://localhost:8080/query?tenant=tenant1
```

### Trigger Model Retraining
```bash
curl http://localhost:8080/retrain?tenant=tenant1
```

---

## Web UI

- Navigate to `http://localhost:8080/`  
- Dashboard shows:
  - Tenant selector  
  - Logs table  
  - Anomaly alerts  
  - Metrics charts (Prometheus/Grafana)  

---

## Helm Deployment

```bash
helm install m2loganalyzer ./helm -n m2-logs
```

- Configure replicas, persistent volumes, and ingress in `values.yaml`  
- Supports multi-tenant isolation and monitoring  

---

## Configuration Example (`configs/config.yaml`)

```yaml
server_port: 8080
workers: 4
queue_size: 1000
tenants:
  - name: tenant1
    useClickHouse: true
    clickHouseDSN: "tcp://localhost:9000?debug=true"
    workers: 4
    queue_size: 1000
  - name: tenant2
    useClickHouse: false
    workers: 2
    queue_size: 500
```

---

## License

Apache 2.0 License – See [LICENSE](LICENSE) file for details.
