package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB   DBConfig
	Auth AuthConfig
}

type DBConfig struct {
	DSN string
}

type AuthConfig struct {
	Secret string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default config")
	}
	return &Config{
		DB: DBConfig{
			DSN: os.Getenv("DSN"),
		},
		Auth: AuthConfig{
			Secret: os.Getenv("SECRET"),
		},
	}

}
