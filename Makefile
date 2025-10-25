.PHONY: build run test clean dev

# ビルド
build:
	go build -o bin/app cmd/app/main.go

# 実行
run: build
	./bin/app

# 開発モード (ホットリロードなし)
dev:
	go run cmd/app/main.go

# テスト
test:
	go test -v ./...

# カバレッジ付きテスト
test-coverage:
	go test -v -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# 依存関係のインストール
deps:
	go mod download
	go mod tidy

# クリーンアップ
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

# フォーマット
fmt:
	go fmt ./...

# リント
lint:
	golangci-lint run

# ヘルプ
help:
	@echo "利用可能なコマンド:"
	@echo "  make build         - アプリケーションをビルド"
	@echo "  make run           - アプリケーションを実行"
	@echo "  make dev           - 開発モードで実行"
	@echo "  make test          - テストを実行"
	@echo "  make test-coverage - カバレッジ付きでテストを実行"
	@echo "  make deps          - 依存関係をインストール"
	@echo "  make clean         - ビルド成果物を削除"
	@echo "  make fmt           - コードをフォーマット"
	@echo "  make lint          - リントを実行"
