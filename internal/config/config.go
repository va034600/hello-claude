package config

import "os"

// Config はアプリケーションの設定を保持します
type Config struct {
	Port string
	Env  string
}

// Load は環境変数から設定を読み込みます
func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	env := os.Getenv("ENV")
	if env == "" {
		env = "development"
	}

	return &Config{
		Port: port,
		Env:  env,
	}
}
