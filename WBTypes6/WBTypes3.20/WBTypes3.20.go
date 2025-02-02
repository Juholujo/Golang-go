package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Пример целого числа, которое мы хотим преобразовать в строку
	num := 12345

	// Преобразование целого числа в строку с помощью strconv.Itoa
	str := strconv.Itoa(num)
	fmt.Printf("Число %d преобразовано в строку: '%s'\n", num, str)

	// Пример целого числа с отрицательным значением
	negativeNum := -9876

	// Преобразование отрицательного числа в строку
	negativeStr := strconv.Itoa(negativeNum)
	fmt.Printf("Отрицательное число %d преобразовано в строку: '%s'\n", negativeNum, negativeStr)

	// Пример использования преобразованного числа в другом контексте
	// Допустим, мы хотим создать сообщение об ошибке с кодом ошибки
	errorCode := 404
	errorMessage := "Error: " + strconv.Itoa(errorCode) + " Not Found"
	fmt.Println(errorMessage)
}
