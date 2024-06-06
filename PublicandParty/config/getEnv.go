package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	TCP_PORT string

	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string

	SERVICE_HOST string
	SERVICE_PORT string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}
	config.TCP_PORT = cast.ToString(coalesce("TCP_PORT", ":50051"))

	config.DB_HOST = cast.ToString(coalesce("DB_HOST", "localhost"))
	config.DB_PORT = cast.ToInt(coalesce("DB_PORT", 5432))
	config.DB_USER = cast.ToString(coalesce("DB_USER", "azizbek"))
	config.DB_PASSWORD = cast.ToString(coalesce("DB_PASSWORD", "123"))
	config.DB_NAME = cast.ToString(coalesce("DB_NAME", "voting2"))

	config.SERVICE_HOST = cast.ToString(coalesce("SERVICE_HOST", "localhost"))
	config.SERVICE_PORT = cast.ToString(coalesce("SERVICE_PORT", ":50051"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
