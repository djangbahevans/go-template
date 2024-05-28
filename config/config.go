package config

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	DBName     string
	DBUser     string
	DBPass     string
	DBHost     string
	DBPort     string
	ServerAddr string
)

func LoadConfig() {
	godotenv.Load()

	DBName = os.Getenv("DB_NAME")
	DBUser = os.Getenv("DB_USER")
	DBPass = os.Getenv("DB_PASS")
	DBHost = os.Getenv("DB_HOST")
	DBPort = os.Getenv("DB_PORT")
	ServerAddr = os.Getenv("SERVER_ADDR")

	// Validate the environment variables
	if ServerAddr == "" {
		ServerAddr = ":8080"
	}
}
