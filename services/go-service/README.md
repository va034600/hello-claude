# Go Service

Ginを使用したシンプルなGo APIサーバーです。

## 概要

このサービスは、Gin Web Frameworkを使用した基本的なREST APIサーバーの実装です。`/api/hello` エンドポイントを提供し、JSON形式でレスポンスを返します。

クリーンアーキテクチャを意識した構成で、ビジネスロジック層とHTTPハンドラー層を分離しています。

## 必要要件

- Go 1.21以上

## インストール

```bash
cd services/go-service
go mod download
```

## 使用方法

### サーバーの起動

```bash
go run main.go
```

サーバーは `http://localhost:3000` で起動します。

### APIエンドポイントのテスト

```bash
curl http://localhost:3000/api/hello
```

レスポンス例:
```json
{"message": "Hello from server!"}
```

### テストの実行

```bash
go test -v ./...
```

## プロジェクト構成

```
go-service/
├── main.go             # エントリーポイント
├── internal/
│   ├── api/            # HTTPハンドラ（ルーティング）
│   │   ├── router.go
│   │   ├── handler.go
│   │   └── handler_test.go
│   ├── service/        # ビジネスロジック層
│   │   ├── hello.go
│   │   └── hello_test.go
│   ├── model/          # ドメインモデル、構造体
│   │   └── response.go
│   └── config/         # 設定読み込み
│       └── config.go
├── go.mod
└── README.md
```
