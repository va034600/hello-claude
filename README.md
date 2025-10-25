# hello-claude

Ginを使用したシンプルなGo APIサーバーのサンプルプロジェクトです。

## 概要

このプロジェクトは、Gin Web Frameworkを使用した基本的なREST APIサーバーの実装例です。`/api/hello` エンドポイントを提供し、JSON形式でレスポンスを返します。

クリーンアーキテクチャを意識した構成で、ビジネスロジック層とHTTPハンドラー層を分離しています。

## 必要要件

- Go 1.21以上

## インストール

1. リポジトリをクローンします:
```bash
git clone https://github.com/va034600/hello-claude.git
cd hello-claude
```

2. 依存パッケージをインストールします:
```bash
make deps
# または
go mod download
```

3. 環境変数を設定します(オプション):
```bash
cp .env.template .env
# .envファイルを編集して必要な値を設定
```

## 使用方法

### サーバーの起動

```bash
make run
# または
make dev  # 開発モード
```

直接実行する場合:
```bash
go run cmd/app/main.go
```

サーバーは `http://localhost:3000` で起動します。

### APIエンドポイントのテスト

サーバー起動後、以下のコマンドでAPIをテストできます:

```bash
curl http://localhost:3000/api/hello
```

レスポンス例:
```json
{"message": "Hello from server!"}
```

### テストの実行

```bash
make test
# または
go test -v ./...
```

カバレッジ付きでテストを実行:
```bash
make test-coverage
```

## 依存関係

- **gin-gonic/gin** (v1.10.0): 高速で柔軟なGo Webフレームワーク

## プロジェクト構成

```
hello-claude/
├── cmd/
│   └── app/            # mainパッケージ、エントリーポイント
│       └── main.go
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
│   └── config/         # 設定読み込み（env, yamlなど）
│       └── config.go
├── pkg/                # 再利用可能な共通ライブラリ
├── go.mod
├── go.sum
├── Makefile           # ビルド・テストコマンド
└── README.md          # このファイル
```

## 利用可能なMakeコマンド

- `make build` - アプリケーションをビルド
- `make run` - アプリケーションを実行
- `make dev` - 開発モードで実行
- `make test` - テストを実行
- `make test-coverage` - カバレッジ付きでテストを実行
- `make deps` - 依存関係をインストール
- `make clean` - ビルド成果物を削除
- `make fmt` - コードをフォーマット
- `make help` - ヘルプを表示

## ライセンス

ISC
