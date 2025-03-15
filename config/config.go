package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	SMTPUser     string
	SMTPPassword string
	DBConnStr    string
)

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	SMTPUser = os.Getenv("SMTP_USER")
	SMTPPassword = os.Getenv("SMTP_PASS")
	DBConnStr = os.Getenv("DATABASE_URL")
}
