package main

import "fmt"

// Универсальная функция для вывода данных
func printValue(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Printf("The value is a string: %s\n", v)
	case int:
		fmt.Printf("The value is an integer: %d\n", v)
	case bool:
		fmt.Printf("The value is a boolean: %t\n", v)
	default:
		fmt.Printf("Unknown type\n")
	}
}

func main() {
	printValue("hello world") // строка
	printValue(42)            // число
	printValue(true)          // булево значение
	printValue(3.14)          // неизвестный тип
}
