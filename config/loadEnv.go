package config

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadEnv() {

	err := godotenv.Load()

	if err != nil {
		envPath := filepath.Join("..", ".env")
		err = godotenv.Load(envPath)
	}

	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}
}
