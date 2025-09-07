# -------- Build Stage --------
FROM golang:1.22-alpine AS builder

WORKDIR /app

# Copy Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy project files
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o m2loganalyzer ./cmd/m2loganalyzer

# -------- Runtime Stage --------
FROM gcr.io/distroless/base

WORKDIR /app

# Copy binary and config
COPY --from=builder /app/m2loganalyzer /app/m2loganalyzer
COPY configs/config.yaml /app/config.yaml

EXPOSE 8080

ENTRYPOINT ["/app/m2loganalyzer"]
