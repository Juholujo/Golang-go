package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/joho/godotenv"
)

func main() {
    // Загружаем .env
    err := godotenv.Load()
    if err != nil {
        log.Println("Нет .env файла или произошла ошибка при загрузке:", err)
    }

    // Считываем переменную PORT
    port := os.Getenv("PORT")
    if port == "" {
        port = ":8080"
        log.Printf("PORT не задан в .env, используется по умолчанию %s", port)
    }

    // Настраиваем простой роут
    http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Server is up"))
    })

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case http.MethodGet:
            w.WriteHeader(http.StatusOK)
            w.Write([]byte("Hello from Go!"))
        case http.MethodPost:
            w.WriteHeader(http.StatusOK)
            w.Write([]byte("Got a POST request"))
        default:
            w.WriteHeader(http.StatusMethodNotAllowed)
        }
    })

    // Запускаем сервер
    log.Printf("Сервер стартует на порту %s ...", port)
    err = http.ListenAndServe(port, nil)
    if err != nil {
        log.Fatalf("Ошибка запуска сервера: %v", err)
    }
}
