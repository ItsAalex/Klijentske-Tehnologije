package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var configurations Config

type Config struct {
	App    App
	Server Server
	DB     DB
}

type App struct {
	Name string
}

type Server struct {
	Host   string
	Port   string
	Secret string
}

type DB struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}

// LoadConfig reads the configuration from environment variables
func LoadConfig() {
	// Load package configuration
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Populate Config struct with environment variables
	configurations.App.Name = getEnv("APP_NAME", "")

	configurations.Server.Host = getEnv("SERVER_HOST", "localhost")
	configurations.Server.Port = getEnv("SERVER_PORT", "8081")
	configurations.Server.Secret = getEnv("SERVER_SECRET", "")

	configurations.DB.Username = getEnv("DB_USERNAME", "root")
	configurations.DB.Password = getEnv("DB_PASSWORD", "root")
	configurations.DB.Host = getEnv("DB_HOST", "172.24.0.2")
	configurations.DB.Port = getEnv("DB_PORT", "3306")
	configurations.DB.Name = getEnv("DB_NAME", "BudoarAtelier")
}

// Get returns the global configurations
func Get() Config {
	return configurations
}

// getEnv is a helper function to get an environment variable with a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	return value
}
