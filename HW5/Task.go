package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
)

// Структура Student содержит данные о студенте.
type Student struct {
    FullName         string `json:"FullName"`
    MathScore        int    `json:"MathScore"`
    InformaticsScore int    `json:"InformaticsScore"`
    EnglishScore     int    `json:"EnglishScore"`
}

// Срез (глобальный) для хранения поступивших студентов
var admittedStudents []Student

// Обработчик для маршрута /apply
func applyHandler(w http.ResponseWriter, r *http.Request) {
    // Разрешаем только POST-запрос
    if r.Method != http.MethodPost {
        http.Error(w, "Только POST-запросы поддерживаются", http.StatusMethodNotAllowed)
        return
    }

    // Декодируем входящий JSON в структуру Student
    var s Student
    err := json.NewDecoder(r.Body).Decode(&s)
    if err != nil {
        http.Error(w, "Ошибка декодирования JSON", http.StatusBadRequest)
        return
    }

    // Проверяем сумму баллов
    totalScore := s.MathScore + s.InformaticsScore + s.EnglishScore
    if totalScore >= 14 {
        // Добавляем в список поступивших
        admittedStudents = append(admittedStudents, s)
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "Студент %s поступил! Сумма баллов: %d\n", s.FullName, totalScore)
    } else {
        // Студент не прошёл порог
        w.WriteHeader(http.StatusOK)
        fmt.Fprintf(w, "Студент %s не поступил. Сумма баллов: %d\n", s.FullName, totalScore)
    }
}

// Обработчик для маршрута /admitted
func admittedHandler(w http.ResponseWriter, r *http.Request) {
    // Разрешаем только GET-запрос
    if r.Method != http.MethodGet {
        http.Error(w, "Только GET-запросы поддерживаются", http.StatusMethodNotAllowed)
        return
    }

    // Возвращаем список поступивших студентов в формате JSON
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(admittedStudents)
}

func main() {
    // Регистрируем обработчики
    http.HandleFunc("/apply", applyHandler)
    http.HandleFunc("/admitted", admittedHandler)

    fmt.Println("Сервер запущен на порту 8080...")
    // Запускаем сервер на порту 8080
    log.Fatal(http.ListenAndServe(":8080", nil))
}
