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

	// Попытка сложить строки (не получится)
	log.Printf("Timeout value is: %s", timeoutStr)
	// Ошибка! Невозможно сложить строки с числами
	// result := timeoutStr + 10 // Ошибка компиляции
	//time.Sleep(time.Duration(timeoutStr) * time.Second) // Применяем число для задержки
	log.Println("Process completed")
}
