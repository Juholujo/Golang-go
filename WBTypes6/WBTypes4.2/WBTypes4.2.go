package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
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

	// Преобразование строки в число
	timeout, err := strconv.Atoi(timeoutStr)
	if err != nil {
		log.Fatalf("Invalid timeout value: %v", err)
	}

	// Пример использования этого значения как числа
	log.Printf("Timeout value is: %d seconds", timeout)

	// Используем значение для таймаута
	log.Println("Starting process with timeout...")
	time.Sleep(time.Duration(timeout) * time.Second) // Применяем число для задержки
	log.Println("Process completed")
}
