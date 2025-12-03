package config

import "os"

type Config struct {
	Port         string
	NodeJSAPIURL string
}

func Load() *Config {
	return &Config{
		Port:         getEnv("PORT", "4000"),
		NodeJSAPIURL: getEnv("NODEJS_API_URL", "http://node-api:5000"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
