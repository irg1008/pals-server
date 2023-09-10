package config

import (
	"fmt"
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

	EmailHost    string
	EmailPort    int
	EmailUser    string
	EmailPass    string
	EmailAddress string

	GoogleClientID     string
	GoogleClientSecret string

	AppleClientID   string
	AppleTeamID     string
	AppleKeyID      string
	ApplePrivateKey string
}

func NewConfig() *Config {
	err := godotenv.Load()

	env := os.Getenv("GO_ENV")
	isDev := env == "development"

	if err != nil && isDev {
		slog.Warn("Failed to load from env file, will try to load from env variables")
	}

	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPass := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := os.Getenv("POSTGRES_HOST")
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPass, dbHost, dbPort, dbName)

	fmt.Print(dbUrl)

	return &Config{
		Env:   env,
		IsDev: isDev,
		Port:  ":" + os.Getenv("PORT"),

		DBUrl:     dbUrl,
		JWTSecret: os.Getenv("JWT_SECRET"),

		ClientUrl: os.Getenv("CLIENT_URL"),
		APIUrl:    os.Getenv("API_URL"),
		Domain:    os.Getenv("DOMAIN"),
		AppName:   os.Getenv("APP_NAME"),

		EmailHost:    os.Getenv("HOST_ADDRESS"),
		EmailPort:    587,
		EmailUser:    os.Getenv("HOST_USERNAME"),
		EmailPass:    os.Getenv("HOST_PASSWORD"),
		EmailAddress: os.Getenv("HOST_ADDRESS"),

		GoogleClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		GoogleClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),

		AppleClientID:   os.Getenv("APPLE_CLIENT_ID"),
		AppleTeamID:     os.Getenv("APPLE_TEAM_ID"),
		AppleKeyID:      os.Getenv("APPLE_KEY_ID"),
		ApplePrivateKey: os.Getenv("APPLE_PRIVATE_KEY"),
	}
}
