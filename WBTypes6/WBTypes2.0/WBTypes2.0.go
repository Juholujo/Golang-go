package main

import "fmt"

func main() {
	// Создаем переменную типа интерфейс и присваиваем ей строку "hello"
	var i interface{} = "0"

	// Пробуем привести интерфейс i к строковому типу

	s := i.(string)

	fmt.Println("Значение s: ", s)

	n := i.(int)

	fmt.Println("Значение n: ", n)

}
