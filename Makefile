.PHONY: help run-go run-node build-go test-go test-node test-all deps-go deps-node deps clean fmt-go

# デフォルトターゲット
.DEFAULT_GOAL := help

# Go Service関連
GO_SERVICE_DIR := services/go-service
GO_BINARY := $(GO_SERVICE_DIR)/server

# Node.js Service関連
NODE_SERVICE_DIR := services/node-service

# ヘルプ
help:
	@echo "利用可能なコマンド:"
	@echo "  make run-go          - Go Serviceを実行"
	@echo "  make run-node        - Node.js Serviceを実行"
	@echo "  make build-go        - Go Serviceをビルド"
	@echo "  make test-go         - Go Serviceのテストを実行"
	@echo "  make test-node       - Node.js Serviceのテストを実行"
	@echo "  make test-all        - 全サービスのテストを実行"
	@echo "  make test-coverage-go - Go Serviceのカバレッジ付きテストを実行"
	@echo "  make deps-go         - Go Serviceの依存関係をインストール"
	@echo "  make deps-node       - Node.js Serviceの依存関係をインストール"
	@echo "  make deps            - 全サービスの依存関係をインストール"
	@echo "  make clean           - ビルド成果物を削除"
	@echo "  make fmt-go          - Go Serviceのコードをフォーマット"

# Go Serviceを実行
run-go:
	@echo "Starting Go service..."
	cd $(GO_SERVICE_DIR) && go run main.go

# Node.js Serviceを実行
run-node:
	@echo "Starting Node.js service..."
	cd $(NODE_SERVICE_DIR) && npm start

# Go Serviceをビルド
build-go:
	@echo "Building Go service..."
	cd $(GO_SERVICE_DIR) && go build -o server main.go

# Go Serviceのテストを実行
test-go:
	@echo "Running Go service tests..."
	cd $(GO_SERVICE_DIR) && go test -v ./...

# Node.js Serviceのテストを実行
test-node:
	@echo "Running Node.js service tests..."
	cd $(NODE_SERVICE_DIR) && npm test

# 全サービスのテストを実行
test-all: test-go test-node

# Go Serviceのカバレッジ付きテストを実行
test-coverage-go:
	@echo "Running Go service tests with coverage..."
	cd $(GO_SERVICE_DIR) && go test -coverprofile=coverage.out ./...
	cd $(GO_SERVICE_DIR) && go tool cover -html=coverage.out -o coverage.html

# Go Serviceの依存関係をインストール
deps-go:
	@echo "Installing Go service dependencies..."
	cd $(GO_SERVICE_DIR) && go mod download
	cd $(GO_SERVICE_DIR) && go mod tidy

# Node.js Serviceの依存関係をインストール
deps-node:
	@echo "Installing Node.js service dependencies..."
	cd $(NODE_SERVICE_DIR) && npm install

# 全サービスの依存関係をインストール
deps: deps-go deps-node

# クリーンアップ
clean:
	@echo "Cleaning build artifacts..."
	rm -f $(GO_BINARY)
	cd $(GO_SERVICE_DIR) && rm -f coverage.out coverage.html
	cd $(NODE_SERVICE_DIR) && rm -rf node_modules

# Go Serviceのコードをフォーマット
fmt-go:
	@echo "Formatting Go code..."
	cd $(GO_SERVICE_DIR) && go fmt ./...
