package main

import (
	"flag"
	"fmt"
	"log"
	"github.com/1lker/sd-gen-o2/internal/api"
	"github.com/1lker/sd-gen-o2/internal/config"
)

func main() {
	// Parse command line flags
	configPath := flag.String("config", "config/config.json", "path to configuration file")
	flag.Parse()

	// Load configuration
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		fmt.Printf("Warning: Could not load config file, using defaults: %v\n", err)
		cfg = config.DefaultConfig()
	}

	fmt.Println("=== Super Dummy Data Generator API ===")
	fmt.Printf("Mode: %s\n", cfg.Server.Mode)
	fmt.Printf("Rate Limit: %d requests per minute\n", cfg.API.RateLimit)

	// Create and start server
	server := api.NewServer(cfg)
	if err := server.Start(); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}