package model

// HelloResponse は /api/hello のレスポンス構造体
type HelloResponse struct {
	Message string `json:"message"`
}

// ErrorResponse はエラーレスポンスの構造体
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
