package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)
var DatabaseURL string
var DatabaseName string
func init() {
	err := godotenv.Load("./config/config.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	DatabaseURL = os.Getenv("DB_URI")
	DatabaseName = os.Getenv("DB_NAME")
}