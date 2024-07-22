package config

import (
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() error {
	return godotenv.Load(".env")
}

func GetEnv(key, defaultValue string) string {
	value, exist := os.LookupEnv(key)
	if !exist {
		return defaultValue
	}
	return value
}
