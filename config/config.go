package config

import (
	"log"
	"time"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
)

func FiberConfig() fiber.Config {
	return fiber.Config{
		ReadTimeout: time.Second * time.Duration(*AppCfg().ReadTimeout),
	}
}

// LoadAllConfigs set various configs
func LoadAllConfigs(envFile string) {

	err := godotenv.Load(envFile)
	if err != nil {
		log.Fatalf("can't load .env file. error: %v", err)
	}

	LoadApp()
}
