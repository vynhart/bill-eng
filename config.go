package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func GetConfig() map[string]string {
	fileName := ".env"
	if os.Getenv("ENV") == "test" {
		fileName = ".env.testing"
	}
	
	err := godotenv.Load(fileName)
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	return map[string]string{
		"db_username": os.Getenv("DB_USERNAME"),
		"db_password": os.Getenv("DB_PASSWORD"),
		"db_host":     os.Getenv("DB_HOST"),
		"db_name":     os.Getenv("DB_NAME"),
	}
}
