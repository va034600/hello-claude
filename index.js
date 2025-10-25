const express = require("express");
const app = express();
const port = 3000;

// 静的ファイルを配信（publicフォルダにindex.htmlなどを置く）
app.use(express.static("public"));

// 例: APIエンドポイント
app.get("/api/hello", (req, res) => {
  res.json({ message: "Hello from server!" });
});

app.listen(port, () => {
  console.log(`Server running at http://localhost:${port}`);
});
