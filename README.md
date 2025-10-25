# hello-claude

マイクロサービスアーキテクチャを採用したシンプルなAPIサーバーのサンプルプロジェクトです。

## 概要

このプロジェクトは、GoとNode.jsの2つのマイクロサービスで構成されています。どちらも同じ `/api/hello` エンドポイントを提供し、JSON形式でレスポンスを返します。

### サービス構成

- **Go Service**: Gin Web Frameworkを使用したGoベースのマイクロサービス
  - クリーンアーキテクチャを採用
  - ポート: 3000（デフォルト）

- **Node.js Service**: Express.jsを使用したNode.jsベースのマイクロサービス
  - シンプルなMVCアーキテクチャ
  - ポート: 3000（デフォルト）

## プロジェクト構成

```
hello-claude/
├── services/
│   ├── go-service/          # Goマイクロサービス
│   │   ├── main.go
│   │   ├── internal/
│   │   │   ├── api/         # HTTPハンドラ
│   │   │   ├── service/     # ビジネスロジック
│   │   │   ├── model/       # データモデル
│   │   │   └── config/      # 設定
│   │   ├── go.mod
│   │   └── README.md
│   │
│   └── node-service/        # Node.jsマイクロサービス
│       ├── index.js
│       ├── index.test.js
│       ├── package.json
│       └── README.md
│
├── Makefile                 # ビルド・実行コマンド
└── README.md               # このファイル
```

## 必要要件

### Go Service
- Go 1.21以上

### Node.js Service
- Node.js 18以上
- npm または yarn

## インストール

### リポジトリをクローン

```bash
git clone https://github.com/va034600/hello-claude.git
cd hello-claude
```

### Go Serviceのセットアップ

```bash
cd services/go-service
go mod download
cd ../..
```

### Node.js Serviceのセットアップ

```bash
cd services/node-service
npm install
cd ../..
```

## 使用方法

### Makefileを使用する場合

プロジェクトルートから各サービスを起動できます:

```bash
# Go Serviceを起動
make run-go

# Node.js Serviceを起動
make run-node

# 両方のサービスをテスト
make test-all
```

### 個別に起動する場合

#### Go Service

```bash
cd services/go-service
go run main.go
```

#### Node.js Service

```bash
cd services/node-service
npm start
```

### APIエンドポイントのテスト

どちらのサービスも同じエンドポイントを提供します（ポートを適宜変更してください）:

```bash
curl http://localhost:3000/api/hello
```

レスポンス例:
```json
{"message": "Hello from server!"}
```

## テスト

### Go Serviceのテスト

```bash
cd services/go-service
go test -v ./...
```

### Node.js Serviceのテスト

```bash
cd services/node-service
npm test
```

### 全サービスのテスト（Makefile使用）

```bash
make test-all
```

## 開発

各サービスは独立して開発・デプロイが可能です。詳細は各サービスのREADMEを参照してください:

- [Go Service README](services/go-service/README.md)
- [Node.js Service README](services/node-service/README.md)

## 利用可能なMakeコマンド

- `make run-go` - Go Serviceを実行
- `make run-node` - Node.js Serviceを実行
- `make test-go` - Go Serviceのテストを実行
- `make test-node` - Node.js Serviceのテストを実行
- `make test-all` - 全サービスのテストを実行
- `make build-go` - Go Serviceをビルド
- `make clean` - ビルド成果物を削除
- `make help` - ヘルプを表示

## ライセンス

ISC
