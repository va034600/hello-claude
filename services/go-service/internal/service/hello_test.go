package service

import "testing"

func TestGetHelloMessage(t *testing.T) {
	service := NewHelloService()

	response := service.GetHelloMessage()

	expected := "Hello from server!"
	if response.Message != expected {
		t.Errorf("期待されるメッセージ: %s, 実際: %s", expected, response.Message)
	}
}
