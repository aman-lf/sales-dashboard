package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

var Cfg Config

type Config struct {
	Port     string
	Host     string
	Mode     string
	MongoURI string
	AppURL   string
}

func init() {
	loadConfig()
}

func loadConfig() {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load(filepath.Join(workingDir, ".env"))

	if err != nil {
		log.Fatal("Error loading .env file, using system environment variables")
	}

	Cfg = Config{
		Port:     os.Getenv("PORT"),
		Host:     os.Getenv("HOST"),
		Mode:     os.Getenv("MODE"),
		MongoURI: os.Getenv("MONGO_URI"),
		AppURL:   os.Getenv("APP_URL"),
	}
}
