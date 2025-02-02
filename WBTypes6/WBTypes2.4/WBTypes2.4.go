package main

import (
	"fmt"
)

// Определение пользовательского типа ошибки
type MyError struct {
	Code int
	Msg  string
}

func (e MyError) Error() string {
	return e.Msg
}

func main() {
	// Создаем ошибку типа MyError
	originalErr := MyError{Code: 404, Msg: "Not Found"}

	// Оборачиваем ошибку с дополнительным контекстом
	wrappedErr := fmt.Errorf("something went wrong: %w", originalErr)

	// Пробуем привести обернутую ошибку к типу MyError
	if myErr, ok := wrappedErr.(MyError); ok {
		fmt.Println("Unwrapped error:", myErr)
	} else {
		fmt.Println("Type assertion failed")
	}
}
