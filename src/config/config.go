package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	OpenAIModel string
	OpenAIApiKey string
	OpenAIOrganization string
	OpenAIProject string
	OpenAIURL string
}

var AppConfig Config

func LoadEnv() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	AppConfig = Config{
		Port: getEnv("PORT", "8080"),
		OpenAIModel: os.Getenv("OPENAI_MODEL"),
		OpenAIApiKey: os.Getenv("OPENAI_API_KEY"),
		OpenAIOrganization: os.Getenv("OPENAI_ORGANIZATION"),
		OpenAIProject: os.Getenv("OPENAI_PROJECT"),
		OpenAIURL: os.Getenv("OPENAI_URL"),
	}
}

func GetPort() string {
	return AppConfig.Port
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
