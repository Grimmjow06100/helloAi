package config

import "os"

type Config struct {
	AppEnv       string
	HTTPAddr     string
	OpenAIAPIKey string
	DatabaseURL  string
	PromptsDir   string
}

func Load() Config {
	return Config{
		AppEnv:       getEnv("APP_ENV", "development"),
		HTTPAddr:     getEnv("HTTP_ADDR", ":8080"),
		OpenAIAPIKey: os.Getenv("OPENAI_API_KEY"),
		DatabaseURL:  os.Getenv("DATABASE_URL"),
		PromptsDir:   getEnv("PROMPTS_DIR", "./prompts"),
	}
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}
