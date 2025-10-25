package service

import "github.com/va034600/hello-claude/services/go-service/internal/model"

// HelloService はhelloエンドポイントのビジネスロジックを提供します
type HelloService struct{}

// NewHelloService は新しいHelloServiceを作成します
func NewHelloService() *HelloService {
	return &HelloService{}
}

// GetHelloMessage はhelloメッセージを返します
func (s *HelloService) GetHelloMessage() *model.HelloResponse {
	return &model.HelloResponse{
		Message: "Hello from server!",
	}
}
