package config

import (
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	Env       string
	Port      string
	IsDev     bool
	DBName    string
	DBUrl     string
	JWTSecret string
}

func LoadEnv() config {
	godotenv.Load()
	env := os.Getenv("GO_ENV")

	return config{
		Env:       env,
		Port:      ":" + os.Getenv("PORT"),
		IsDev:     env == "development",
		DBName:    os.Getenv("DB_NAME"),
		DBUrl:     os.Getenv("DB_URL"),
		JWTSecret: os.Getenv("JWT_SECRET"),
	}
}

var Env config

func init() {
	Env = LoadEnv()
}
