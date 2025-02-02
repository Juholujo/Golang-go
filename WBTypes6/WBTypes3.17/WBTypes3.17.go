package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Пример строки, которую мы хотим преобразовать в int
	str := "12345"

	// Преобразование строки в int с помощью strconv.Atoi
	if num, err := strconv.Atoi(str); err == nil {
		fmt.Printf("String '%s' преобразована в int: %d\n", str, num)
	} else {
		fmt.Printf("Ошибка при преобразовании строки '%s': %v\n", str, err)
	}

	// Пример строки, которая не может быть преобразована в int
	invalidStr := "abc"

	// Обработка ошибки при попытке преобразования неправильной строки
	if num, err := strconv.Atoi(invalidStr); err == nil {
		fmt.Printf("String '%s' преобразована в int: %d\n", invalidStr, num)
	} else {
		fmt.Printf("Ошибка при преобразовании строки '%s': %v\n", invalidStr, err)
	}
}
