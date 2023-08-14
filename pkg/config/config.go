package config

import (
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
}

func NewConfig() *Config {
	godotenv.Load()
	env := os.Getenv("GO_ENV")

	return &Config{
		Env:       env,
		Port:      ":" + os.Getenv("PORT"),
		IsDev:     env == "development",
		DBName:    os.Getenv("DB_NAME"),
		DBUrl:     os.Getenv("DB_URL"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		ClientUrl: os.Getenv("CLIENT_URL"),
	}
}
