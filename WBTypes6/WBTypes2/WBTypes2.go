package main

import "fmt"

func main() {
	// Создаем переменную типа интерфейс и присваиваем ей строку "hello"
	var i interface{} = "0"

	// Пробуем привести интерфейс i к строковому типу
	s, ok := i.(string)

	//s := i.(string)

	fmt.Println("Значение ok: ", ok)

	fmt.Println("Значение s: ", s)

	// ok будет true, если приведение удалось, и false, если нет
	if ok {
		fmt.Println("The value is a string:", s)
	} else {
		fmt.Println("The value is not a string")
	}

	// Пробуем привести интерфейс i к типу int
	n, ok := i.(int)
	//n := i.(int)

	fmt.Println("Значение ok: ", ok)
	fmt.Println("Значение n: ", n)

	// Здесь ok будет false, потому что i содержит строку, а не int
	if ok {
		fmt.Println("The value is an int:", n)
	} else {
		fmt.Println("The value is not an int")
	}
}
