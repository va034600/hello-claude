# hello-claude

Express.jsを使用したシンプルなNode.js APIサーバーのサンプルプロジェクトです。

## 概要

このプロジェクトは、Express 5.1.0を使用した基本的なREST APIサーバーの実装例です。`/api/hello` エンドポイントを提供し、JSON形式でレスポンスを返します。

## 必要要件

- Node.js (v14以上推奨)
- npm (Node.jsに含まれています)

## インストール

1. リポジトリをクローンします:
```bash
git clone https://github.com/va034600/hello-claude.git
cd hello-claude
```

2. 依存パッケージをインストールします:
```bash
npm install
```

3. 環境変数を設定します（オプション）:
```bash
cp .env.template .env
# .envファイルを編集して必要な値を設定
```

## 使用方法

### サーバーの起動

```bash
npm start
```

または開発モードで起動:

```bash
npm run dev
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

## 依存関係

- **express** (^5.1.0): 高速で柔軟なNode.js Webアプリケーションフレームワーク

## プロジェクト構成

```
hello-claude/
├── index.js              # メインのサーバーファイル
├── package.json          # プロジェクト設定と依存関係
├── .env.template         # 環境変数のテンプレート
├── .gitignore           # Gitで無視するファイルの設定
└── README.md            # このファイル
```

## ライセンス

ISC
