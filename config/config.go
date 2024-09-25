package config

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
}

func SetUp() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, err
	}

	port := os.Getenv("PORT")
	if port == "" {
		return nil, errors.New("port is empty")
	}

	return &Config{
		Port: port,
	}, nil
}
