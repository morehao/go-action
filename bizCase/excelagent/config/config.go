package config

import (
	"fmt"
	"os"
)

type Config struct {
	Server ServerConfig `yaml:"server"`
	Eino   EinoConfig   `yaml:"eino"`
	Task   TaskConfig   `yaml:"task"`
}

type ServerConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type EinoConfig struct {
	APIKey  string `yaml:"api_key"`
	BaseURL string `yaml:"base_url"`
	Model   string `yaml:"model"`
}

type TaskConfig struct {
	MaxIterations int    `yaml:"max_iterations"`
	Timeout       int    `yaml:"timeout"`
	Workspace     string `yaml:"workspace"`
}

func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Host: getEnv("SERVER_HOST", "0.0.0.0"),
			Port: getEnvInt("SERVER_PORT", 8080),
		},
		Eino: EinoConfig{
			APIKey:  getEnv("EINO_API_KEY", ""),
			BaseURL: getEnv("EINO_BASE_URL", "https://ark.cn-beijing.volces.com/api/v3"),
			Model:   getEnv("EINO_MODEL", "deepseek-v3"),
		},
		Task: TaskConfig{
			MaxIterations: getEnvInt("TASK_MAX_ITERATIONS", 20),
			Timeout:       getEnvInt("TASK_TIMEOUT", 300),
			Workspace:     getEnv("TASK_WORKSPACE", "./workspace"),
		},
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		var intValue int
		if _, err := fmt.Sscanf(value, "%d", &intValue); err == nil {
			return intValue
		}
	}
	return defaultValue
}
