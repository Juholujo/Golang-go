package main

import (
	"errors"
	"fmt"
)

// MyError Определение пользовательского типа ошибки
type MyError struct {
	Code int
	Msg  string
}

func (e MyError) Error() string {
	return e.Msg
}

// AnotherError Другой пользовательский тип ошибки
type AnotherError struct {
	Description string
}

func (e AnotherError) Error() string {
	return e.Description
}

// Функция, которая возвращает разные типы ошибок
func doSomething(flag int) error {
	if flag == 1 {
		return MyError{Code: 404, Msg: "Not Found"}
	} else if flag == 2 {
		return AnotherError{Description: "Something went wrong"}
	} else if flag == 3 {
		return errors.New("a standard error")
	}
	return nil
}

func main() {
	// Вызов функции с разными типами ошибок
	for _, flag := range []int{1, 2, 3, 4} {
		err := doSomething(flag)
		if err != nil {
			// Пробуем привести к MyError
			if myErr, ok := err.(MyError); ok {
				fmt.Println("Custom Error (MyError):")
				fmt.Println("  Code:", myErr.Code)
				fmt.Println("  Message:", myErr.Msg)
			} else if anotherErr, ok := err.(AnotherError); ok {
				fmt.Println("Custom Error (AnotherError):")
				fmt.Println("  Description:", anotherErr.Description)
			} else {
				// Обработка стандартной ошибки
				fmt.Println("Generic Error:", err)
			}
		} else {
			fmt.Println("No error occurred")
		}
	}
}
