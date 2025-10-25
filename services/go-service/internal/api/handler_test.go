package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/va034600/hello-claude/services/go-service/internal/model"
)

func TestGetHello(t *testing.T) {
	// テストモードに設定
	gin.SetMode(gin.TestMode)

	// ルーターの初期化
	router := NewRouter()

	t.Run("正常にレスポンスを返す", func(t *testing.T) {
		// リクエストの作成
		req, _ := http.NewRequest("GET", "/api/hello", nil)
		w := httptest.NewRecorder()

		// リクエストの実行
		router.ServeHTTP(w, req)

		// ステータスコードの確認
		if w.Code != http.StatusOK {
			t.Errorf("期待されるステータスコード: %d, 実際: %d", http.StatusOK, w.Code)
		}

		// レスポンスボディの確認
		var response model.HelloResponse
		if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Fatalf("JSONのパースに失敗: %v", err)
		}

		expected := "Hello from server!"
		if response.Message != expected {
			t.Errorf("期待されるメッセージ: %s, 実際: %s", expected, response.Message)
		}
	})

	t.Run("Content-Typeがapplication/jsonである", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/api/hello", nil)
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		contentType := w.Header().Get("Content-Type")
		if contentType != "application/json; charset=utf-8" {
			t.Errorf("期待されるContent-Type: application/json; charset=utf-8, 実際: %s", contentType)
		}
	})
}

func TestNotFound(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := NewRouter()

	req, _ := http.NewRequest("GET", "/not/found", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("期待されるステータスコード: %d, 実際: %d", http.StatusNotFound, w.Code)
	}

	var response model.ErrorResponse
	if err := json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Fatalf("JSONのパースに失敗: %v", err)
	}

	if response.Error != "Not Found" {
		t.Errorf("期待されるエラー: Not Found, 実際: %s", response.Error)
	}
}
