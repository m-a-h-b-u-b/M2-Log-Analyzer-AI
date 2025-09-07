APP_NAME = m2loganalyzer
DOCKER_IMAGE = m2loganalyzer:latest

.PHONY: all build run docker-build docker-run clean test

all: build

# Build Go binary locally
build:
	@echo "==> Building Go binary..."
	go build -o bin/$(APP_NAME) ./cmd/m2loganalyzer

# Run locally
run: build
	@echo "==> Running locally..."
	./bin/$(APP_NAME)

# Build Docker image
docker-build:
	@echo "==> Building Docker image..."
	docker build -t $(DOCKER_IMAGE) -f deploy/docker/Dockerfile .

# Run Docker container
docker-run: docker-build
	@echo "==> Running Docker container..."
	docker run -p 8080:8080 $(DOCKER_IMAGE)

# Run Go tests
test:
	@echo "==> Running tests..."
	go test ./... -v

# Clean up binaries
clean:
	@echo "==> Cleaning..."
	rm -rf bin
