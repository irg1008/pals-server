package config

import (
	"log/slog"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const (
	TokenDuration        = 15 * time.Minute
	RefreshTokenDuration = 48 * time.Hour
)

type Config struct {
	Env       string
	Port      string
	IsDev     bool
	DBName    string
	DBUrl     string
	JWTSecret string
	ClientUrl string
	ResendKey string
	Domain    string
}

func NewConfig() *Config {
	err := godotenv.Load()

	env := os.Getenv("GO_ENV")
	isDev := env == "development"

	if err != nil && isDev {
		slog.Warn("Failed to load from env file, will try to load from env variables")
	}

	return &Config{
		Env:       env,
		IsDev:     isDev,
		Port:      ":" + os.Getenv("PORT"),
		DBName:    os.Getenv("DB_NAME"),
		DBUrl:     os.Getenv("DB_URL"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		ClientUrl: os.Getenv("CLIENT_URL"),
		ResendKey: os.Getenv("RESEND_KEY"),
		Domain:    os.Getenv("DOMAIN"),
	}
}
