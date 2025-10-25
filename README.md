# hello-claude

マイクロサービスアーキテクチャを採用したシンプルなAPIサーバーのサンプルプロジェクトです。

## 概要

このプロジェクトは、GoとNode.jsの2つのマイクロサービスで構成されています。どちらも同じ `/api/hello` エンドポイントを提供し、JSON形式でレスポンスを返します。

## サービス

- **[Go Service](services/go-service/README.md)** - Gin Web Frameworkを使用したGoベースのマイクロサービス
- **[Node.js Service](services/node-service/README.md)** - Express.jsを使用したNode.jsベースのマイクロサービス

各サービスの詳細なセットアップ方法、使用方法、テスト方法については、それぞれのREADMEを参照してください。

## クイックスタート

### Makefileを使用する場合

```bash
# Go Serviceを起動
make run-go

# Node.js Serviceを起動
make run-node

# 全サービスのテストを実行
make test-all
```

### 利用可能なMakeコマンド

- `make run-go` / `make run-node` - 各サービスを実行
- `make test-go` / `make test-node` - 各サービスのテストを実行
- `make test-all` - 全サービスのテストを実行
- `make build-go` - Go Serviceをビルド
- `make clean` - ビルド成果物を削除
- `make help` - ヘルプを表示

## ライセンス

ISC
