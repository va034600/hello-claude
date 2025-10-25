package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/va034600/hello-claude/internal/service"
)

// HelloHandler は /api/hello エンドポイントのハンドラー
type HelloHandler struct {
	service *service.HelloService
}

// NewHelloHandler は新しいHelloHandlerを作成します
func NewHelloHandler(service *service.HelloService) *HelloHandler {
	return &HelloHandler{
		service: service,
	}
}

// GetHello は GET /api/hello のハンドラー
func (h *HelloHandler) GetHello(c *gin.Context) {
	response := h.service.GetHelloMessage()
	c.JSON(http.StatusOK, response)
}
