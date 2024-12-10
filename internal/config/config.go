package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Server ServerConfig `json:"server"`
	API    APIConfig    `json:"api"`
}

type ServerConfig struct {
	Port            int    `json:"port"`
	Host            string `json:"host"`
	Mode            string `json:"mode"`
	AllowedOrigins  string `json:"allowed_origins"`
}

type APIConfig struct {
	RateLimit      int    `json:"rate_limit"`       // Requests per minute
	TimeoutSeconds int    `json:"timeout_seconds"`
	MaxRequestSize int64  `json:"max_request_size"` // In bytes
}

// DefaultConfig returns default configuration
func DefaultConfig() *Config {
	return &Config{
		Server: ServerConfig{
			Port:           8080,
			Host:           "localhost",
			Mode:           "release",
			AllowedOrigins: "*",
		},
		API: APIConfig{
			RateLimit:      60,    // 60 requests per minute
			TimeoutSeconds: 30,    // 30 seconds timeout
			MaxRequestSize: 1 << 20, // 1MB
		},
	}
}

// LoadConfig loads configuration from a JSON file
func LoadConfig(path string) (*Config, error) {
	config := DefaultConfig()

	file, err := os.Open(path)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		return config, err
	}

	return config, nil
}

// SaveConfig saves configuration to a JSON file
func (c *Config) SaveConfig(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(c)
}