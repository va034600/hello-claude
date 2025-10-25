# Node.js Service

Express.jsを使用したシンプルなNode.js APIサーバーです。

## 概要

このサービスは、Express.jsを使用した基本的なREST APIサーバーの実装です。`/api/hello` エンドポイントを提供し、JSON形式でレスポンスを返します。

## 必要要件

- Node.js 18以上
- npm または yarn

## インストール

```bash
cd services/node-service
npm install
```

## 使用方法

### サーバーの起動

```bash
npm start
# または開発モード
npm run dev
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
npm test
```

## プロジェクト構成

```
node-service/
├── index.js          # メインアプリケーションファイル
├── index.test.js     # テストファイル
├── package.json      # 依存関係とスクリプト
└── README.md         # このファイル
```
