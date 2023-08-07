package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env   string
	Port  string
	IsDev bool
}

func GetConfig() Config {
	godotenv.Load()

	env := os.Getenv("ENV")
	port := ":" + os.Getenv("PORT")

	return Config{
		Env:   env,
		Port:  port,
		IsDev: env == "development",
	}
}
