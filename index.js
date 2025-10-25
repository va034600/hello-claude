const express = require("express");
const app = express();
const port = process.env.PORT || 3000;

// 例: APIエンドポイント
app.get("/api/hello", (req, res) => {
  res.json({ message: "Hello from server!" });
});

// テスト時にはサーバーを起動しない
if (require.main === module) {
  app.listen(port, () => {
    console.log(`Server running at http://localhost:${port}`);
  });
}

module.exports = app;
