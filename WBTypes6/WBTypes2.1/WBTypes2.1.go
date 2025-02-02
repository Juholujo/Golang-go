package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// JSON-строка
	data := `{"name": "Alice", "age": 30, "isAdmin": true}`

	// Декодируем JSON в map[string]interface{}
	var result map[string]interface{}
	err := json.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Обрабатываем данные
	for key, value := range result {
		// Используем type assertion для определения типа
		switch v := value.(type) {
		case string:
			fmt.Printf("%s is a string: %s\n", key, v)
		case float64:
			fmt.Printf("%s is a number: %f\n", key, v)
		case bool:
			fmt.Printf("%s is a boolean: %t\n", key, v)
		default:
			fmt.Printf("%s is of an unknown type\n", key)
		}
	}
}
