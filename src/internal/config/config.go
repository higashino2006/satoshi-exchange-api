package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	FRONTEND_URL          string
	BACKEND_URL           string
	DB_HOST               string
	DB_DATABASE           string
	DB_USERNAME           string
	DB_PASSWORD           string
	MIGRATION_FOLDER_PATH string
}

var AppConfig Config

func Init() error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
		return err
	}

	AppConfig = Config{
		FRONTEND_URL:          os.Getenv("FRONTEND_URL"),
		BACKEND_URL:           os.Getenv("BACKEND_URL"),
		DB_HOST:               os.Getenv("DB_HOST"),
		DB_DATABASE:           os.Getenv("DB_DATABASE"),
		DB_USERNAME:           os.Getenv("DB_USERNAME"),
		DB_PASSWORD:           os.Getenv("DB_PASSWORD"),
		MIGRATION_FOLDER_PATH: os.Getenv("MIGRATION_FOLDER_PATH"),
	}

	return nil
}
