package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Config структура для маппинга данных из YAML
type Config struct {
	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	} `yaml:"app"`

	Server struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"server"`

	Features struct {
		EnableLogging bool `yaml:"enableLogging"`
		MaxRetries    int  `yaml:"maxRetries"`
	} `yaml:"features"`
}

func main() {
	// Открытие YAML файла
	file, err := os.Open("config.yml")
	if err != nil {
		log.Fatalf("Error opening YAML file: %v", err)
	}
	defer file.Close()

	// Парсинг YAML файла
	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("Error decoding YAML file: %v", err)
	}

	// Используем данные из конфигурации
	fmt.Printf("App Name: %s\n", config.App.Name)
	fmt.Printf("App Version: %s\n", config.App.Version)
	fmt.Printf("Server Host: %s\n", config.Server.Host)
	fmt.Printf("Server Port: %d\n", config.Server.Port)
	fmt.Printf("Logging Enabled: %t\n", config.Features.EnableLogging)
	fmt.Printf("Max Retries: %d\n", config.Features.MaxRetries)
}
