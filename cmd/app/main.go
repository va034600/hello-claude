package main

import (
	"log"
	"os"

	"github.com/va034600/hello-claude/internal/api"
	"github.com/va034600/hello-claude/internal/config"
)

func main() {
	// 設定の読み込み
	cfg := config.Load()

	// ルーターの初期化
	router := api.NewRouter()

	// サーバーの起動
	addr := ":" + cfg.Port
	log.Printf("Server running at http://localhost%s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
		os.Exit(1)
	}
}
