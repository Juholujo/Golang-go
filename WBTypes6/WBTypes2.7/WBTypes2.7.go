package main

import "fmt"

func main() {
	// Конфигурация в виде динамической структуры
	config := map[string]interface{}{
		"timeout": 30,
		"debug":   true,
		"host":    "localhost",
	}

	// Извлечение и использование значений
	if timeout, ok := config["timeout"].(int); ok {
		fmt.Println("Timeout is:", timeout)
	} else {
		fmt.Println("Timeout is not an integer")
	}

	if debug, ok := config["debug"].(bool); ok {
		fmt.Println("Debug mode is:", debug)
	} else {
		fmt.Println("Debug mode is not a boolean")
	}

	if host, ok := config["host"].(string); ok {
		fmt.Println("Host is:", host)
	} else {
		fmt.Println("Host is not a string")
	}

	// Попытка извлечь значение, которого нет
	if port, ok := config["port"].(int); ok {
		fmt.Println("Port is:", port)
	} else {
		fmt.Println("Port is not defined or not an integer")
	}
}
