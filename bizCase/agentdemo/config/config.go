package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// AppConfig holds all application configuration.
type AppConfig struct {
	Model  ModelConfig  `yaml:"model"`
	Server ServerConfig `yaml:"server"`
}

// ModelConfig holds LLM model configuration.
type ModelConfig struct {
	APIKey  string `yaml:"api_key"`
	Model   string `yaml:"model"`
	BaseURL string `yaml:"base_url"`
}

// ServerConfig holds HTTP server configuration.
type ServerConfig struct {
	Port string `yaml:"port"`
}

// Cfg is the global application config, pre-loaded with sensible defaults.
var Cfg = &AppConfig{
	Model: ModelConfig{
		Model:   "deepseek-chat",
		BaseURL: "https://api.deepseek.com/v1",
	},
	Server: ServerConfig{
		Port: ":8080",
	},
}

// Load reads the YAML config file at the given path and merges it into Cfg.
// Errors are printed to stderr so the operator is aware of any config issues.
func Load(configPath string) {
	data, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "agentdemo: config: read %s: %v\n", configPath, err)
		return
	}
	if err = yaml.Unmarshal(data, Cfg); err != nil {
		fmt.Fprintf(os.Stderr, "agentdemo: config: parse %s: %v\n", configPath, err)
	}
}
