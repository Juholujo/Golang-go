package main

import (
	"fmt"
	"strings"
)

func main() {
	// 1. Сравнение строк без учета регистра
	str1 := "GoLang"
	str2 := "golang"
	fmt.Println("Сравнение строк без учета регистра:", strings.EqualFold(str1, str2)) // true

	// 2. Поиск подстроки
	str := "hello, world"
	substr := "world"
	fmt.Println("Поиск подстроки:", strings.Contains(str, substr)) // true

	// 3. Подсчет количества вхождений подстроки
	fmt.Println("Подсчет количества вхождений подстроки:", strings.Count(str, "l")) // 3

	// 4. Проверка начинается ли строка с определенной подстроки
	fmt.Println("Проверка начинается ли строка с определенной подстроки:", strings.HasPrefix(str, "hello")) // true

	// 5. Проверка заканчивается ли строка определенной подстрокой
	fmt.Println("Проверка заканчивается ли строка определенной подстрокой:", strings.HasSuffix(str, "world")) // true

	// 6. Возвращает индекс первого вхождения подстроки
	fmt.Println("Возвращает индекс первого вхождения подстроки:", strings.Index(str, "world")) // 7

	// 7. Возвращает индекс последнего вхождения подстроки
	fmt.Println("Возвращает индекс последнего вхождения подстроки:", strings.LastIndex(str, "l")) // 10

	// 8. Обрезает пробельные символы в начале и в конце строки
	fmt.Println("Обрезает пробельные символы в начале и в конце строки:", strings.TrimSpace("     he ll o, world     ")) // "hello, world"

	// 9. Преобразует строку в нижний регистр
	fmt.Println("Преобразует строку в нижний регистр:", strings.ToLower("GoLang")) // "golang"

	// 10. Преобразует строку в верхний регистр
	fmt.Println("Преобразует строку в верхний регистр:", strings.ToUpper("golang")) // "GOLANG"

	// 11. Заменяет все вхождения одной подстроки на другую
	fmt.Println("Заменяет все вхождения одной подстроки на другую:", strings.Replace("hello, hello, hello", "hello", "world", 2)) // "world, world, hello"

	// 12. Разбивает строку на подстроки по разделителю
	fmt.Println("Разбивает строку на подстроки по разделителю:", strings.Split("a,b,c", ",")) // ["a" "b" "c"]

	// 13. Склеивает срез строк в одну строку, разделенную разделителем
	fmt.Println("Склеивает срез строк в одну строку, разделенную разделителем:", strings.Join([]string{"hello", "world"}, ", ")) // "hello, world"

	// 14. Возвращает строку, состоящую из символов, повторенных до указанной длины
	fmt.Println("Возвращает строку, состоящую из символов, повторенных до указанной длины:", strings.Repeat("Go", 3)) // "GoGoGo"

	// 15. Удаляет все вхождения указанных символов в начале и в конце строки
	fmt.Println("Удаляет все вхождения указанных символов в начале и в конце строки:", strings.Trim("! hello! ", "! ")) // "hello"

	// 16. Переворачивает строку
	fmt.Println("Переворачивает строку:", strings.Join(strings.Fields("hello, world"), " ")) // "world hello"

}
