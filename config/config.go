package config

import (
	"os"
)

var Config appConfig

type appConfig struct {
	DBUser     string
	DBPassword string
	DBPort     string
	DBHost     string
	DBName     string
	ServerPort string
}

func LoadConfig() {
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "3306"
	}
	Config = appConfig{
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBPort:     dbPort,
		DBHost:     os.Getenv("DB_HOST"),
		DBName:     os.Getenv("DB_NAME"),
		ServerPort: serverPort,
	}
}
