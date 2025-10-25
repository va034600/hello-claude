package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/va034600/hello-claude/services/go-service/internal/model"
	"github.com/va034600/hello-claude/services/go-service/internal/service"
)

// NewRouter は新しいGinルーターを作成します
func NewRouter() *gin.Engine {
	router := gin.Default()

	// サービスの初期化
	helloService := service.NewHelloService()

	// ハンドラーの初期化
	helloHandler := NewHelloHandler(helloService)

	// ルーティング設定
	api := router.Group("/api")
	{
		api.GET("/hello", helloHandler.GetHello)
	}

	// 404ハンドラー
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, model.ErrorResponse{
			Error:   "Not Found",
			Message: "リクエストされたパス " + c.Request.URL.Path + " は見つかりませんでした",
		})
	})

	return router
}
