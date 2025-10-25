const express = require("express");
const app = express();
const port = process.env.PORT || 3000;

// 例: APIエンドポイント
app.get("/api/hello", (req, res) => {
  res.json({ message: "Hello from server!" });
});

// 404エラーハンドラー - 存在しないルートへのリクエストを処理
app.use((req, res) => {
  res.status(404).json({
    error: "Not Found",
    message: `リクエストされたパス ${req.path} は見つかりませんでした`
  });
});

// グローバルエラーハンドラー - アプリケーション全体のエラーをキャッチ
app.use((err, req, res, next) => {
  console.error('エラーが発生しました:', err.stack);
  res.status(err.status || 500).json({
    error: "Internal Server Error",
    message: process.env.NODE_ENV === 'production'
      ? 'サーバーエラーが発生しました'
      : err.message
  });
});

// テスト時にはサーバーを起動しない
if (require.main === module) {
  app.listen(port, () => {
    console.log(`Server running at http://localhost:${port}`);
  });
}

module.exports = app;
