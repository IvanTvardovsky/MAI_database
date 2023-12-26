package config

import (
	"backend/internal/schemas"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

var ProjectConfig schemas.Config

func createConfig() schemas.Config {
	return schemas.Config{
		DB: schemas.DatabaseConfig{
			User:     "postgres",
			Password: "bebra123",
			Host:     "localhost",
			Port:     "5432",
			DBName:   "MAI_db",
		},
		Deploy: schemas.Deploy{
			Port: 5050,
		},
	}
}

func getStrEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getIntEnv(key string, defaultVal int) int {
	if value, exists := os.LookupEnv(key); exists {
		intVal, _ := strconv.Atoi(value)
		return intVal
	}

	return defaultVal
}

func init() {
	godotenv.Load()
	ProjectConfig = createConfig()
}
