package configs

import (
	"log"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	Db DbConfig
	Auth AuthConfig
}

type DbConfig struct{
	DSN string
}

type AuthConfig struct{
	Secret string
}

func LoadConfig() *Config{
	err := godotenv.Load()
	
	if err != nil {
		log.Println("Error loading")
	}

	return &Config{
		Db: DbConfig{
			DSN: os.Getenv("DSN"),
		},

		Auth: AuthConfig{
			Secret: os.Getenv("SECRET"),
		},
	}
}