package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	// Загрузка переменных окружения из файла .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Получение значения переменной окружения "TIMEOUT"
	timeoutStr := os.Getenv("TIMEOUT")
	if timeoutStr == "" {
		log.Fatal("TIMEOUT environment variable is not set")
	}

	// Здесь timeoutStr остается строкой
	log.Printf("Timeout value is: %s", timeoutStr) // Просто выводим строку
}
