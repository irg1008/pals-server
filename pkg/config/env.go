package config

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env   string
	IsDev bool
	Port  string

	DBUrl     string
	JWTSecret string

	ClientUrl string
	APIUrl    string
	Domain    string
	AppName   string

	GoogleClientID     string
	GoogleClientSecret string

	EmailHost    string
	EmailPort    int
	EmailUser    string
	EmailPass    string
	EmailAddress string

	AuthCoreUrl string
	AuthCoreKey string
}

func NewConfig() *Config {
	err := godotenv.Load()

	env := os.Getenv("GO_ENV")
	isDev := env == "development"

	if err != nil && isDev {
		slog.Warn("Failed to load from env file, will try to load from env variables")
	}

	return &Config{
		Env:                env,
		IsDev:              isDev,
		Port:               ":" + os.Getenv("PORT"),
		JWTSecret:          os.Getenv("JWT_SECRET"),
		ClientUrl:          os.Getenv("CLIENT_URL"),
		APIUrl:             os.Getenv("API_URL"),
		AppName:            os.Getenv("APP_NAME"),
		EmailHost:          os.Getenv("HOST_ADDRESS"),
		EmailPort:          587,
		EmailUser:          os.Getenv("HOST_USERNAME"),
		EmailPass:          os.Getenv("HOST_PASSWORD"),
		EmailAddress:       os.Getenv("HOST_ADDRESS"),
		Domain:             os.Getenv("DOMAIN"),
		AuthCoreUrl:        os.Getenv("AUTH_CORE_URL"),
		AuthCoreKey:        os.Getenv("AUTH_CORE_KEY"),
		DBUrl:              os.Getenv("DB_URL"),
		GoogleClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
	}
}
