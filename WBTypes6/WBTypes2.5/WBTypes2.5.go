package main

import (
	"errors"
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

	// Используем errors.As для извлечения MyError
	var myErr MyError
	if errors.As(wrappedErr, &myErr) {
		fmt.Println("Unwrapped error:")
		fmt.Println("  Code:", myErr.Code)
		fmt.Println("  Message:", myErr.Msg)
	} else {
		fmt.Println("Could not unwrap to MyError")
	}
}

/*Как работает fmt.Errorf("some context: %w", err)?
Когда вы вызываете fmt.Errorf с форматом %w,
Go создает новую ошибку, которая оборачивает (или "wraps") вашу оригинальную ошибку.
Эта новая ошибка является экземпляром внутреннего типа *fmt.wrapError.
Внутри она хранит ссылку на оригинальную ошибку.

Упрощенная структура обернутой ошибки:

type wrapError struct {
	msg string    // Сообщение обертки (например, "some context")
	err error     // Ссылка на оригинальную ошибку (например, MyError)
}*/

/*Иллюстрация цепочки оберток

wrappedErr (тип *fmt.wrapError)
└── originalErr (тип MyError)*/
