package config

import (
	"errors"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// Load env file
func GoDotEnvVariable(key string) (*string, error) {
	err := godotenv.Load(".env")

	if err != nil {
		return nil, errors.New("Error loading .env file")
	}

	loaded := os.Getenv(key)

	return &loaded, nil
}

func LoadAllConfig() {
	LoadApp()
}

func FiberConfig() fiber.Config {
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(*AppCfg().ReadTimeout),
	}
}
