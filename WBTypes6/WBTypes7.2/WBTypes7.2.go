package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
)

// Config структура для хранения данных
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
	// Создаем данные для кодирования
	config := Config{
		App: struct {
			Name    string `yaml:"name"`
			Version string `yaml:"version"`
		}{
			Name:    "MyApp",
			Version: "1.0.0",
		},
		Server: struct {
			Host string `yaml:"host"`
			Port int    `yaml:"port"`
		}{
			Host: "localhost",
			Port: 8080,
		},
		Features: struct {
			EnableLogging bool `yaml:"enableLogging"`
			MaxRetries    int  `yaml:"maxRetries"`
		}{
			EnableLogging: true,
			MaxRetries:    5,
		},
	}

	// Кодируем структуру в YAML-строку
	data, err := yaml.Marshal(&config)
	if err != nil {
		panic(err)
	}

	// Выводим YAML-строку
	fmt.Println(string(data))
}
