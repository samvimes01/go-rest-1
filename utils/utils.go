package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetValue(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file") // Fatal and Panic will break the app in prod mode, but prod env vars shouldn't be in .env
	}
	return os.Getenv(key)
}